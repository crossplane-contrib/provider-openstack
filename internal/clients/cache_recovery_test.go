package clients

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/gophercloud/gophercloud/v2"
	osprovider "github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"
)

func cacheRecoveryRequest(t *testing.T, ctx context.Context, client *gophercloud.ProviderClient, endpoint string) error {
	t.Helper()
	response, err := client.Request(ctx, http.MethodGet, endpoint, &gophercloud.RequestOpts{})
	if response != nil {
		if closeErr := response.Body.Close(); closeErr != nil {
			t.Errorf("close service response: %v", closeErr)
		}
	}
	return err
}

func TestNoForkClientCacheInvalidatesRevokedToken(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	keystone := newKeystoneTestServer(t, time.Hour, nil)
	keystone.onRequest = func(w http.ResponseWriter, r *http.Request) bool {
		if r.URL.Path != "/revoked" {
			return false
		}
		w.WriteHeader(http.StatusUnauthorized)
		return true
	}
	cache := newNoForkClientCache(1)
	provider := osprovider.Provider()
	configuration := keystoneTestConfiguration(keystone)
	configuration["allow_reauth"] = "false"
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	first := terraform.Setup{Configuration: configuration}
	if err := cache.configure(ctx, &first, provider); err != nil {
		t.Fatal(err)
	}
	old := authenticatedOpenstackTestConfig(t, first)
	if err := cacheRecoveryRequest(t, ctx, old.OsClient, keystone.server.URL+"/revoked"); err == nil {
		t.Fatal("expected an authentication rejection")
	}
	if got := keystone.tokenRequests.Load(); got != 1 {
		t.Fatalf("allow_reauth=false made an automatic token request: %d", got)
	}
	second := terraform.Setup{Configuration: configuration}
	if err := cache.configure(ctx, &second, provider); err != nil {
		t.Fatal(err)
	}
	if second.Meta == first.Meta || keystone.tokenRequests.Load() != 2 {
		t.Fatal("a revoked token remained cached until its scheduled expiry")
	}
	// An old asynchronous operation can reject the old token after the
	// replacement is already cached. It must not evict that replacement.
	if err := cacheRecoveryRequest(t, ctx, old.OsClient, keystone.server.URL+"/revoked"); err == nil {
		t.Fatal("expected the old token to remain rejected")
	}
	third := terraform.Setup{Configuration: configuration}
	if err := cache.configure(ctx, &third, provider); err != nil {
		t.Fatal(err)
	}
	if third.Meta != second.Meta || keystone.tokenRequests.Load() != 2 {
		t.Fatal("a late rejection invalidated the replacement generation")
	}
}

func TestNoForkClientCacheAutomaticReauthentication(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	keystone := newKeystoneTestServer(t, time.Hour, nil)
	keystone.onRequest = func(w http.ResponseWriter, r *http.Request) bool {
		if r.URL.Path != "/service" {
			return false
		}
		if r.Header.Get("X-Auth-Token") == "dummy-token-1" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		return true
	}
	cache := newNoForkClientCache(1)
	provider := osprovider.Provider()
	setup := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	if err := cache.configure(ctx, &setup, provider); err != nil {
		t.Fatal(err)
	}
	cfg := authenticatedOpenstackTestConfig(t, setup)
	if err := cacheRecoveryRequest(t, ctx, cfg.OsClient, keystone.server.URL+"/service"); err != nil {
		t.Fatalf("automatic reauthentication failed: %v", err)
	}
	if cfg.OsClient.Token() != "dummy-token-2" || keystone.tokenRequests.Load() != 2 {
		t.Fatal("the existing client did not automatically refresh its rejected token")
	}
}

func TestNoForkClientCacheReauthenticationTimeout(t *testing.T) {
	if withReauthTimeout(nil, time.Second) != nil {
		t.Fatal("a client without automatic reauthentication must remain unchanged")
	}
	t.Run("bounded", func(t *testing.T) {
		wrapped := withReauthTimeout(func(ctx context.Context) error {
			deadline, ok := ctx.Deadline()
			if !ok || time.Until(deadline) > time.Millisecond {
				t.Error("reauthentication has no bounded deadline")
			}
			<-ctx.Done()
			return ctx.Err()
		}, time.Millisecond)
		if err := wrapped(t.Context()); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("expected a bounded authentication timeout, got %v", err)
		}
	})
	t.Run("canceled parent", func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		wrapped := withReauthTimeout(func(ctx context.Context) error { return ctx.Err() }, time.Minute)
		if err := wrapped(ctx); !errors.Is(err, context.Canceled) {
			t.Fatalf("authentication ignored parent cancellation: %v", err)
		}
	})
}
