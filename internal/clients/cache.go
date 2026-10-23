package clients

import (
	"container/list"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	tokens2 "github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	osprovider "github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"
	"golang.org/x/sync/singleflight"
)

const (
	defaultNoForkClientCacheSize = 128
	noForkTokenExpiryClockSkew   = 5 * time.Minute
	noForkReauthTimeout          = 30 * time.Second
)

var errNoForkInitializationCanceled = errors.New("OpenStack client initialization canceled")

// noForkClientCache shares authenticated provider Meta values across reconciles
// with identical configurations. Only successful Keystone authentication with a
// known token expiry is cached; failures and Swift authentication stay uncached.
// A separate configured provider copy is still needed for each configuration so
// that credentials from different ProviderConfigs cannot mix.
type noForkClientCache struct {
	maxSize int
	now     func() time.Time
	mu      sync.Mutex
	entries map[string]*list.Element
	// order contains exactly one element per entry, oldest first. Expiration
	// and eviction remove entries from both structures under the same lock.
	order    list.List
	inflight singleflight.Group
}

type noForkClientCacheEntry struct {
	key     string
	meta    any
	expires time.Time
}

func newNoForkClientCache(maxSize int) *noForkClientCache {
	if maxSize <= 0 {
		maxSize = defaultNoForkClientCacheSize
	}
	return &noForkClientCache{
		maxSize: maxSize,
		now:     time.Now,
		entries: make(map[string]*list.Element),
	}
}

// getOrCreate coalesces simultaneous misses for one configuration without
// blocking initialization of other configurations. A waiting reconcile can
// cancel independently. The creator must not mutate the caller's Setup because
// creation may finish after that caller has canceled.
func (c *noForkClientCache) getOrCreate(ctx context.Context, config map[string]any, create func() (any, error)) (any, error) {
	key, err := hashNoForkClientConfig(config)
	if err != nil {
		return nil, err
	}
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if meta, ok := c.get(key); ok {
			return meta, nil
		}

		result := c.inflight.DoChan(key, func() (any, error) {
			return c.initialize(ctx, key, create)
		})
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case r := <-result:
			// A live caller must not inherit the initializing caller's canceled
			// context. The failed flight is already removed, so retrying here
			// coalesces a fresh initialization under a live caller's context.
			if errors.Is(r.Err, errNoForkInitializationCanceled) {
				continue
			}
			return r.Val, r.Err
		}
	}
}

func (c *noForkClientCache) initialize(ctx context.Context, key string, create func() (any, error)) (any, error) {
	// Another creator may have populated the cache since the first check.
	if meta, ok := c.get(key); ok {
		return meta, nil
	}
	meta, err := create()
	if err != nil {
		if ctx.Err() != nil {
			return nil, errors.Join(errNoForkInitializationCanceled, ctx.Err(), err)
		}
		return nil, err
	}
	if _, ok := tokenExpiry(meta, c.now()); ok {
		cfg := meta.(*osprovider.Config)
		base := cfg.OsClient.HTTPClient.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		// Install hooks before publication. A revoked token must not remain
		// sticky until its scheduled expiry, even with allow_reauth=false.
		cfg.OsClient.HTTPClient.Transport = &noForkClientTransport{
			base:       base,
			invalidate: func() { c.invalidate(key, cfg) },
		}
		cfg.OsClient.ReauthFunc = withReauthTimeout(cfg.OsClient.ReauthFunc, noForkReauthTimeout)
	}
	c.set(key, meta)
	return meta, nil
}

func (c *noForkClientCache) get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	entry := element.Value.(*noForkClientCacheEntry)
	if !c.now().Before(entry.expires) {
		c.remove(element)
		return nil, false
	}
	return entry.meta, true
}

func (c *noForkClientCache) set(key string, meta any) {
	expires, ok := tokenExpiry(meta, c.now())
	if !ok {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	entry := &noForkClientCacheEntry{key: key, meta: meta, expires: expires}
	if element, ok := c.entries[key]; ok {
		element.Value = entry
		return
	}
	if len(c.entries) >= c.maxSize {
		c.remove(c.order.Front())
	}
	c.entries[key] = c.order.PushBack(entry)
}

// remove requires c.mu to be held.
func (c *noForkClientCache) remove(element *list.Element) {
	delete(c.entries, element.Value.(*noForkClientCacheEntry).key)
	c.order.Remove(element)
}

func (c *noForkClientCache) invalidate(key string, meta *osprovider.Config) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// In-flight operations may still use an evicted client. A late 401 from
	// that generation must not invalidate its replacement.
	if element, ok := c.entries[key]; ok && element.Value.(*noForkClientCacheEntry).meta == meta {
		c.remove(element)
	}
}

type noForkClientTransport struct {
	base       http.RoundTripper
	invalidate func()
}

func (t *noForkClientTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := t.base.RoundTrip(request)
	if response != nil && response.StatusCode == http.StatusUnauthorized {
		t.invalidate()
	}
	return response, err
}

// Gophercloud's shared reauthentication wait does not observe each waiting
// request's context. Bound the authentication itself so a stalled Keystone
// request cannot hold other reconciles until an async operation's hour deadline.
func withReauthTimeout(reauth func(context.Context) error, timeout time.Duration) func(context.Context) error {
	if reauth == nil {
		return nil
	}
	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return reauth(ctx)
	}
}

// JSON orders map keys deterministically. Hash the whole provider configuration
// so credential rotation or a change in scope/region never reuses an old client.
// Only the digest is retained as the key; credentials are never logged.
func hashNoForkClientConfig(config map[string]any) (string, error) {
	data, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

// tokenExpiry uses the token obtained before cache insertion. Without a known
// expiry we cannot safely reuse a client, particularly with allow_reauth=false.
func tokenExpiry(meta any, now time.Time) (time.Time, bool) {
	cfg, ok := meta.(*osprovider.Config)
	if !ok || cfg == nil || cfg.Swauth || cfg.OsClient == nil || cfg.OsClient.Token() == "" {
		return time.Time{}, false
	}

	expiresAt := keystoneTokenExpiry(cfg.OsClient.GetAuthResult())
	remaining := expiresAt.Sub(now)
	if remaining <= 0 {
		return time.Time{}, false
	}
	// A fixed five-minute margin would disable caching for short-lived tokens.
	// Reserve ten percent of their remaining lifetime, capped at five minutes.
	margin := min(noForkTokenExpiryClockSkew, remaining/10)
	return expiresAt.Add(-margin), true
}

func keystoneTokenExpiry(result any) time.Time {
	switch result := result.(type) {
	case interface {
		ExtractToken() (*tokens3.Token, error)
	}:
		token, err := result.ExtractToken()
		if err == nil && token != nil {
			return token.ExpiresAt
		}
	case interface {
		ExtractToken() (*tokens2.Token, error)
	}:
		token, err := result.ExtractToken()
		if err == nil && token != nil {
			return token.ExpiresAt
		}
	}
	return time.Time{}
}
