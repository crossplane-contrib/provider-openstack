package clients

import (
	"sync"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	osprovider "github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"
)

func TestNoForkClientCacheSwiftNotShared(t *testing.T) {
	for _, fromEnvironment := range []bool{false, true} {
		name := "explicit"
		if fromEnvironment {
			name = "environment default"
		}
		t.Run(name, func(t *testing.T) {
			configuration := configFor("dummy-password")
			if fromEnvironment {
				t.Setenv("OS_SWAUTH", "true")
			} else {
				configuration["swauth"] = "true"
			}
			// The unsafe non-delayed Swift path must also bypass sharing.
			configuration["delayed_auth"] = "false"
			cache := newNoForkClientCache(1)
			provider := osprovider.Provider()
			const callers = 10
			metas := make(chan *osprovider.Config, callers)
			var wg sync.WaitGroup
			for range callers {
				wg.Add(1)
				go func() {
					defer wg.Done()
					setup := terraform.Setup{Configuration: configuration}
					if err := cache.configure(t.Context(), &setup, provider); err != nil {
						t.Errorf("Swift configuration failed: %v", err)
						return
					}
					meta, ok := setup.Meta.(*osprovider.Config)
					if !ok || !meta.Swauth {
						t.Error("expected a Swift configuration")
						return
					}
					metas <- meta
				}()
			}
			wg.Wait()
			close(metas)
			seen := make(map[*osprovider.Config]bool)
			for meta := range metas {
				if seen[meta] {
					t.Error("concurrent callers shared a Swift configuration")
				}
				seen[meta] = true
			}
			if len(seen) != callers || len(cache.entries) != 0 || cache.order.Len() != 0 {
				t.Fatalf("Swift clients=%d, cache entries=%d, eviction records=%d", len(seen), len(cache.entries), cache.order.Len())
			}
		})
	}
}
