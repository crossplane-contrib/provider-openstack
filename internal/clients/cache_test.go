package clients

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	tokens2 "github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	osprovider "github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"
)

func configFor(secret string) map[string]any {
	return map[string]any{
		"auth_url":  "https://identity.example.com/v3",
		"user_name": "crossplane",
		"password":  secret,
		"insecure":  "false",
	}
}

func cacheTestKey(t *testing.T, config map[string]any) string {
	t.Helper()
	key, err := hashNoForkClientConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	return key
}

func cacheTestClient(t *testing.T, expires time.Time) *osprovider.Config {
	t.Helper()
	result := tokens3.CreateResult{}
	result.Header = http.Header{"X-Subject-Token": {"dummy-token"}}
	result.Body = map[string]any{"token": map[string]any{"expires_at": expires.Format(time.RFC3339Nano)}}
	return cacheTestClientWithResult(t, result)
}

func cacheTestClientWithResult(t *testing.T, result gophercloud.AuthResult) *osprovider.Config {
	t.Helper()
	client, err := openstack.NewClient("https://identity.example.com/v3")
	if err != nil {
		t.Fatal(err)
	}
	if err := client.SetTokenAndAuthResult(result); err != nil {
		t.Fatal(err)
	}
	cfg := &osprovider.Config{}
	cfg.OsClient = client
	return cfg
}

func TestNoForkClientCacheHashConfig(t *testing.T) {
	a := cacheTestKey(t, map[string]any{"auth_url": "https://identity.example.com/v3", "user_name": "crossplane"})
	b := cacheTestKey(t, map[string]any{"user_name": "crossplane", "auth_url": "https://identity.example.com/v3"})
	if a != b {
		t.Fatal("identical configurations must hash identically")
	}
	for _, field := range []string{"password", "tenant_id", "region", "auth_url", "insecure"} {
		t.Run(field, func(t *testing.T) {
			config := configFor("dummy-password")
			before := cacheTestKey(t, config)
			config[field] = "changed"
			if before == cacheTestKey(t, config) {
				t.Fatal("changed configuration reused the same key")
			}
		})
	}
}

func TestNoForkClientCacheGetSet(t *testing.T) {
	cache := newNoForkClientCache(2)
	key := cacheTestKey(t, configFor("dummy-password"))
	if _, ok := cache.get(key); ok {
		t.Fatal("expected an initial cache miss")
	}
	meta := cacheTestClient(t, time.Now().Add(time.Hour))
	cache.set(key, meta)
	if got, ok := cache.get(key); !ok || got != meta {
		t.Fatal("expected the original authenticated client")
	}
	if _, ok := cache.get(cacheTestKey(t, configFor("rotated-password"))); ok {
		t.Fatal("rotated credentials must not reuse the previous client")
	}
}

func TestNoForkClientCacheExpirationAndEviction(t *testing.T) {
	now := time.Now()
	cache := newNoForkClientCache(1)
	cache.now = func() time.Time { return now }

	// Reproduce stale FIFO records: expire A, insert B, then insert C.
	cache.set("A", cacheTestClient(t, now.Add(time.Minute)))
	now = now.Add(time.Minute)
	if _, ok := cache.get("A"); ok {
		t.Fatal("expired entry was returned")
	}
	if len(cache.entries) != 0 || cache.order.Len() != 0 {
		t.Fatal("expiration left stale eviction records")
	}
	cache.set("B", cacheTestClient(t, now.Add(time.Hour)))
	cache.set("C", cacheTestClient(t, now.Add(time.Hour)))
	if _, ok := cache.get("B"); ok {
		t.Fatal("oldest live entry was not evicted")
	}
	if _, ok := cache.get("C"); !ok {
		t.Fatal("newest entry was evicted")
	}

	// Repeated expiration/reinsertion must not grow either structure.
	for range 300 {
		now = now.Add(2 * time.Hour)
		cache.get("C")
		cache.set("C", cacheTestClient(t, now.Add(time.Hour)))
		if len(cache.entries) != 1 || cache.order.Len() != 1 {
			t.Fatalf("cache exceeded capacity: map=%d, order=%d", len(cache.entries), cache.order.Len())
		}
	}
}

func TestNoForkClientCacheReplaceEntry(t *testing.T) {
	cache := newNoForkClientCache(2)
	expires := time.Now().Add(time.Hour)
	cache.set("A", cacheTestClient(t, expires))
	cache.set("B", cacheTestClient(t, expires))
	replacement := cacheTestClient(t, expires)
	cache.set("A", replacement)
	if got, ok := cache.get("A"); !ok || got != replacement {
		t.Fatal("replacement entry was not returned")
	}
	if len(cache.entries) != 2 || cache.order.Len() != 2 {
		t.Fatal("replacement changed the number of entries")
	}
	cache.set("C", cacheTestClient(t, expires))
	if _, ok := cache.get("A"); ok {
		t.Fatal("replacement must not duplicate or reorder FIFO records")
	}
}

func TestNoForkClientCacheUncacheable(t *testing.T) {
	now := time.Now()
	swift := cacheTestClient(t, now.Add(time.Hour))
	swift.Swauth = true
	clientWithoutToken, err := openstack.NewClient("https://identity.example.com/v3")
	if err != nil {
		t.Fatal(err)
	}
	unauthenticated := &osprovider.Config{}
	unauthenticated.OsClient = clientWithoutToken
	for _, tt := range []struct {
		name string
		meta any
	}{
		{name: "nil"},
		{name: "typed nil", meta: (*osprovider.Config)(nil)},
		{name: "unknown meta", meta: "not-a-config"},
		{name: "missing client", meta: &osprovider.Config{}},
		{name: "unauthenticated", meta: unauthenticated},
		{name: "missing expiry", meta: cacheTestClient(t, time.Time{})},
		{name: "expired token", meta: cacheTestClient(t, now.Add(-time.Second))},
		{name: "Swift", meta: swift},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cache := newNoForkClientCache(1)
			cache.now = func() time.Time { return now }
			cache.set("key", tt.meta)
			if len(cache.entries) != 0 || cache.order.Len() != 0 {
				t.Fatal("unsafe client was cached")
			}
		})
	}
}

func TestNoForkClientCacheTokenExpiry(t *testing.T) {
	now := time.Now().UTC()
	for _, lifetime := range []time.Duration{time.Hour, 10 * time.Minute, time.Minute} {
		t.Run(lifetime.String(), func(t *testing.T) {
			meta := cacheTestClient(t, now.Add(lifetime))
			expires, ok := tokenExpiry(meta, now)
			want := now.Add(lifetime - min(noForkTokenExpiryClockSkew, lifetime/10))
			if !ok || !expires.Equal(want) {
				t.Fatalf("expiry = %v, %t; want %v", expires, ok, want)
			}
			cache := newNoForkClientCache(1)
			cache.now = func() time.Time { return now }
			cache.set("key", meta)
			now = expires
			if _, ok := cache.get("key"); ok {
				t.Fatal("client was returned at its cache expiry boundary")
			}
		})
	}
}

func TestNoForkClientCacheV2TokenExpiry(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Millisecond)
	result := tokens2.CreateResult{Result: gophercloud.Result{Body: map[string]any{
		"access": map[string]any{"token": map[string]any{
			"id": "dummy-token", "expires": now.Add(time.Hour).Format(gophercloud.RFC3339Milli),
		}},
	}}}
	meta := cacheTestClientWithResult(t, result)
	if expires, ok := tokenExpiry(meta, now); !ok || !expires.Equal(now.Add(55*time.Minute)) {
		t.Fatalf("Keystone v2 expiry = %v, %t", expires, ok)
	}
}

func TestNoForkClientCacheConcurrentCreation(t *testing.T) {
	cache := newNoForkClientCache(1)
	meta := cacheTestClient(t, time.Now().Add(time.Hour))
	var calls atomic.Int32
	var wg sync.WaitGroup
	start := make(chan struct{})
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			got, err := cache.getOrCreate(t.Context(), configFor("dummy-password"), func() (any, error) {
				calls.Add(1)
				return meta, nil
			})
			if err != nil || got != meta {
				t.Errorf("getOrCreate = %v, %v", got, err)
			}
		}()
	}
	close(start)
	wg.Wait()
	if got := calls.Load(); got != 1 {
		t.Fatalf("creator called %d times, want 1", got)
	}
}

func TestNoForkClientCacheFailedCreationRetries(t *testing.T) {
	cache := newNoForkClientCache(1)
	wantErr := errors.New("temporary authentication failure")
	if _, err := cache.getOrCreate(t.Context(), configFor("dummy-password"), func() (any, error) {
		return nil, wantErr
	}); !errors.Is(err, wantErr) {
		t.Fatalf("expected authentication error, got %v", err)
	}
	meta := cacheTestClient(t, time.Now().Add(time.Hour))
	got, err := cache.getOrCreate(t.Context(), configFor("dummy-password"), func() (any, error) {
		return meta, nil
	})
	if err != nil || got != meta {
		t.Fatalf("retry did not recover immediately: %v, %v", got, err)
	}
}

func TestNoForkClientCacheCanceledWaiter(t *testing.T) {
	cache := newNoForkClientCache(2)
	meta := cacheTestClient(t, time.Now().Add(time.Hour))
	started, release, done := make(chan struct{}), make(chan struct{}), make(chan struct{})
	var releaseOnce sync.Once
	unblock := func() { releaseOnce.Do(func() { close(release) }) }
	go func() {
		defer close(done)
		_, err := cache.getOrCreate(t.Context(), configFor("dummy-password"), func() (any, error) {
			close(started)
			<-release
			return meta, nil
		})
		if err != nil {
			t.Errorf("leader failed: %v", err)
		}
	}()
	defer func() {
		unblock()
		<-done
	}()
	<-started
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if _, err := cache.getOrCreate(ctx, configFor("dummy-password"), func() (any, error) {
		t.Error("canceled waiter invoked the creator")
		return nil, nil
	}); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
	// An unrelated key must initialize even while the first key is blocked.
	if _, err := cache.getOrCreate(t.Context(), configFor("other-password"), func() (any, error) {
		return meta, nil
	}); err != nil {
		t.Fatal(err)
	}
	// Release and wait before the test ends so no background writes escape it.
	// The deferred unblock also releases the leader if an assertion fails above.
	unblock()
	<-done
}
