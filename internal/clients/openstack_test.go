package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/gophercloud/gophercloud/v2"
	tokens3 "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	osprovider "github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	clustercomputev1alpha1 "github.com/crossplane-contrib/provider-openstack/apis/cluster/compute/v1alpha1"
	clusterv1beta1 "github.com/crossplane-contrib/provider-openstack/apis/cluster/v1beta1"
	namespacedcomputev1alpha1 "github.com/crossplane-contrib/provider-openstack/apis/namespaced/compute/v1alpha1"
	namespacedv1beta1 "github.com/crossplane-contrib/provider-openstack/apis/namespaced/v1beta1"
)

// These tests intentionally use the upstream provider's real configuration and
// gophercloud authentication paths. In particular, Config.Authenticate remembers
// its first failure, so caching an unauthenticated Config poisons later setups.
type keystoneTestServer struct {
	server        *httptest.Server
	tokenRequests atomic.Int32
	lifetime      time.Duration
	now           func() time.Time
	onToken       func(http.ResponseWriter, *http.Request, int32) bool
	onRequest     func(http.ResponseWriter, *http.Request) bool
	mu            sync.Mutex
	credentials   []keystoneTestCredentials
}

type keystoneTestCredentials struct {
	password string
	tenantID string
}

func newKeystoneTestServer(t *testing.T, lifetime time.Duration, onToken func(http.ResponseWriter, *http.Request, int32) bool) *keystoneTestServer {
	t.Helper()
	k := &keystoneTestServer{lifetime: lifetime, now: time.Now, onToken: onToken}
	k.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/":
			// Exercise version discovery, rather than replacing gophercloud's
			// authenticator or supplying a preconfigured ProviderClient.
			if err := json.NewEncoder(w).Encode(map[string]any{
				"versions": map[string]any{"values": []any{map[string]any{
					"id": "v3.14", "status": "stable",
					"links": []any{map[string]any{"rel": "self", "href": k.server.URL + "/v3/"}},
				}}},
			}); err != nil {
				t.Errorf("encode Keystone discovery: %v", err)
			}
		case r.Method == http.MethodPost && r.URL.Path == "/v3/auth/tokens":
			attempt := k.tokenRequests.Add(1)
			var request struct {
				Auth struct {
					Identity struct {
						Password struct {
							User struct {
								Name     string `json:"name"`
								Password string `json:"password"`
							} `json:"user"`
						} `json:"password"`
					} `json:"identity"`
					Scope struct {
						Project struct {
							ID string `json:"id"`
						} `json:"project"`
					} `json:"scope"`
				} `json:"auth"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Errorf("decode Keystone authentication: %v", err)
				http.Error(w, "invalid authentication JSON", http.StatusBadRequest)
				return
			}
			if request.Auth.Identity.Password.User.Name != "dummy-user" || request.Auth.Identity.Password.User.Password == "" || request.Auth.Scope.Project.ID == "" {
				t.Error("Keystone request did not contain the configured user, password, and project scope")
				http.Error(w, "missing credentials", http.StatusUnauthorized)
				return
			}
			k.mu.Lock()
			k.credentials = append(k.credentials, keystoneTestCredentials{
				password: request.Auth.Identity.Password.User.Password,
				tenantID: request.Auth.Scope.Project.ID,
			})
			k.mu.Unlock()
			if k.onToken != nil && k.onToken(w, r, attempt) {
				return
			}
			now := k.now().UTC()
			w.Header().Set("X-Subject-Token", fmt.Sprintf("dummy-token-%d", attempt))
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(map[string]any{
				"token": map[string]any{
					"methods":    []string{"password"},
					"issued_at":  now.Format(time.RFC3339),
					"expires_at": now.Add(k.lifetime).Format(time.RFC3339),
					"user":       map[string]any{"id": "dummy-user-id", "name": "dummy-user"},
					"project":    map[string]any{"id": request.Auth.Scope.Project.ID, "name": "dummy-project"},
					"catalog": []any{map[string]any{
						"id": "dummy-identity", "name": "keystone", "type": "identity",
						"endpoints": []any{map[string]any{
							"id": "dummy-endpoint", "interface": "public", "region": "dummy-region", "url": k.server.URL + "/v3/",
						}},
					}},
				},
			}); err != nil {
				t.Errorf("encode Keystone token: %v", err)
			}
		default:
			if k.onRequest != nil && k.onRequest(w, r) {
				return
			}
			t.Errorf("unexpected Keystone request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(k.server.Close)
	return k
}

func clearOpenstackTestEnvironment(t *testing.T) {
	t.Helper()
	// Do not let a developer's openrc or CI credentials change provider defaults
	// or direct a test at anything other than its local Keystone fixture.
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		if strings.HasPrefix(name, "OS_") {
			t.Setenv(name, "")
		}
	}
	t.Setenv("TF_LOG", "")
}

func keystoneTestConfiguration(k *keystoneTestServer) map[string]any {
	return map[string]any{
		"auth_url":       k.server.URL,
		"user_name":      "dummy-user",
		"password":       "dummy-password",
		"tenant_id":      "dummy-project",
		"user_domain_id": "default",
		"region":         "dummy-region",
		"insecure":       "false",
		"max_retries":    "0",
	}
}

func authenticatedOpenstackTestConfig(t *testing.T, ps terraform.Setup) *osprovider.Config {
	t.Helper()
	cfg, ok := ps.Meta.(*osprovider.Config)
	if !ok || cfg == nil || cfg.OsClient == nil {
		t.Fatalf("expected an OpenStack Config with a provider client, got %T", ps.Meta)
	}
	if cfg.OsClient.Token() == "" {
		t.Fatal("setup returned successfully without authenticating its provider client")
	}
	result, ok := cfg.OsClient.GetAuthResult().(tokens3.CreateResult)
	if !ok {
		t.Fatalf("expected a real Keystone v3 authentication result, got %T", cfg.OsClient.GetAuthResult())
	}
	if _, err := result.ExtractToken(); err != nil {
		t.Fatalf("extract authenticated Keystone token: %v", err)
	}
	return cfg
}

func TestConfigureNoForkOpenstackClientAuthenticates(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	for _, tc := range []struct {
		name        string
		delayedAuth string
		wantDelayed bool
	}{
		{name: "DefaultDelayedAuth", wantDelayed: true},
		{name: "EagerAuth", delayedAuth: "false"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			keystone := newKeystoneTestServer(t, time.Hour, nil)
			provider := osprovider.Provider()
			ps := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
			if tc.delayedAuth != "" {
				ps.Configuration["delayed_auth"] = tc.delayedAuth
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := configureNoForkOpenstackClient(ctx, &ps, provider); err != nil {
				t.Fatalf("configure: %v", err)
			}
			cfg := authenticatedOpenstackTestConfig(t, ps)
			if cfg.DelayedAuth != tc.wantDelayed {
				t.Errorf("DelayedAuth = %t, want %t", cfg.DelayedAuth, tc.wantDelayed)
			}
			if got := keystone.tokenRequests.Load(); got != 1 {
				t.Errorf("token requests after configure = %d, want 1", got)
			}
			if err := cfg.Authenticate(ctx); err != nil {
				t.Fatalf("authenticate already configured client: %v", err)
			}
			if got := keystone.tokenRequests.Load(); got != 1 {
				t.Errorf("authenticating an already configured client requested %d tokens, want 1", got)
			}
			endpoint, err := cfg.OsClient.EndpointLocator(gophercloud.EndpointOpts{
				Type: "identity", Region: "dummy-region", Availability: gophercloud.AvailabilityPublic,
			})
			if err != nil || endpoint != keystone.server.URL+"/v3/" {
				t.Errorf("authenticated catalog lookup = (%q, %v), want local identity endpoint", endpoint, err)
			}
			if provider.Meta() != nil {
				t.Error("configuration mutated the shared provider singleton's Meta")
			}
		})
	}
}

func TestNoForkClientCacheConfigureCoalescesAuthentication(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	for _, delayedAuth := range []string{"default", "false"} {
		t.Run(delayedAuth, func(t *testing.T) {
			tokenStarted := make(chan struct{})
			releaseToken := make(chan struct{})
			keystone := newKeystoneTestServer(t, time.Hour, func(_ http.ResponseWriter, r *http.Request, attempt int32) bool {
				if attempt == 1 {
					close(tokenStarted)
					select {
					case <-releaseToken:
					case <-r.Context().Done():
						return true
					}
				}
				return false
			})
			var releaseOnce sync.Once
			release := func() { releaseOnce.Do(func() { close(releaseToken) }) }
			t.Cleanup(release)
			provider := osprovider.Provider()
			cache := newNoForkClientCache(8)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			const callers = 16
			start := make(chan struct{})
			entered := make(chan struct{}, callers)
			type result struct {
				setup terraform.Setup
				err   error
			}
			results := make(chan result, callers)
			for range callers {
				go func() {
					<-start
					ps := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
					if delayedAuth != "default" {
						ps.Configuration["delayed_auth"] = delayedAuth
					}
					entered <- struct{}{}
					err := cache.configure(ctx, &ps, provider)
					results <- result{setup: ps, err: err}
				}()
			}
			close(start)
			for range callers {
				select {
				case <-entered:
				case <-ctx.Done():
					t.Fatal("timed out starting concurrent setups")
				}
			}
			select {
			case <-tokenStarted:
			case <-ctx.Done():
				t.Fatal("concurrent setups did not reach Keystone authentication")
			}
			release()
			var shared *osprovider.Config
			for range callers {
				select {
				case result := <-results:
					if result.err != nil {
						t.Fatalf("concurrent configure: %v", result.err)
					}
					cfg := authenticatedOpenstackTestConfig(t, result.setup)
					if shared == nil {
						shared = cfg
					} else if cfg != shared {
						t.Error("concurrent setups did not share the same authenticated Config")
					}
				case <-ctx.Done():
					t.Fatal("timed out waiting for concurrent setups")
				}
			}
			ps := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
			if delayedAuth != "default" {
				ps.Configuration["delayed_auth"] = delayedAuth
			}
			if err := cache.configure(ctx, &ps, provider); err != nil {
				t.Fatalf("subsequent configure: %v", err)
			}
			if ps.Meta != shared || keystone.tokenRequests.Load() != 1 {
				t.Errorf("expected one shared Config and token across concurrent and subsequent setups; token requests = %d", keystone.tokenRequests.Load())
			}
			if provider.Meta() != nil {
				t.Error("concurrent configuration mutated the shared provider singleton's Meta")
			}
		})
	}
}

func TestNoForkClientCacheConfigureRetriesInitialRateLimit(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	keystone := newKeystoneTestServer(t, time.Hour, func(w http.ResponseWriter, _ *http.Request, attempt int32) bool {
		if attempt != 1 {
			return false
		}
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"error":{"code":429,"title":"Too Many Requests","message":"dummy rate limit"}}`))
		return true
	})
	provider := osprovider.Provider()
	cache := newNoForkClientCache(8)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	first := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
	if err := cache.configure(ctx, &first, provider); err == nil {
		t.Fatal("expected the initial Keystone 429 to fail setup")
	}
	if got := keystone.tokenRequests.Load(); got != 1 {
		t.Fatalf("initial failed setup requested %d tokens, want 1", got)
	}
	// No clock advance, backoff, or eviction: a failed delayed authentication
	// must not leave its sticky authFailed Config in the cache.
	retry := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
	if err := cache.configure(ctx, &retry, provider); err != nil {
		t.Fatalf("immediate retry after initial 429: %v", err)
	}
	shared := authenticatedOpenstackTestConfig(t, retry)
	if err := shared.Authenticate(ctx); err != nil {
		t.Fatalf("retry retained a failed authentication: %v", err)
	}
	third := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
	if err := cache.configure(ctx, &third, provider); err != nil {
		t.Fatalf("configure after successful retry: %v", err)
	}
	if third.Meta != shared || keystone.tokenRequests.Load() != 2 {
		t.Errorf("successful retry was not reused; token requests = %d, want 2", keystone.tokenRequests.Load())
	}
	if provider.Meta() != nil {
		t.Error("failed or retried configuration mutated the provider singleton's Meta")
	}
}

func TestNoForkClientCacheConfigureDoesNotCacheCanceledAuthentication(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	tokenStarted := make(chan struct{})
	releaseToken := make(chan struct{})
	keystone := newKeystoneTestServer(t, time.Hour, func(_ http.ResponseWriter, r *http.Request, attempt int32) bool {
		if attempt != 1 {
			return false
		}
		close(tokenStarted)
		select {
		case <-r.Context().Done():
		case <-releaseToken:
		}
		return true
	})
	t.Cleanup(func() { close(releaseToken) })
	provider := osprovider.Provider()
	cache := newNoForkClientCache(8)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	first := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
	done := make(chan error, 1)
	go func() { done <- cache.configure(ctx, &first, provider) }()
	select {
	case <-tokenStarted:
	case <-time.After(5 * time.Second):
		t.Fatal("setup did not reach initial authentication")
	}
	cancel()
	select {
	case err := <-done:
		// The upstream HTTP retry transport formats its underlying error with
		// %s, so cancellation is not necessarily retained in the error chain.
		if err == nil || (!errors.Is(err, context.Canceled) && !strings.Contains(err.Error(), context.Canceled.Error())) {
			t.Fatalf("canceled authentication error = %v, want context.Canceled", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("initial authentication did not honor context cancellation")
	}
	retryCtx, retryCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer retryCancel()
	retry := terraform.Setup{Configuration: keystoneTestConfiguration(keystone)}
	if err := cache.configure(retryCtx, &retry, provider); err != nil {
		t.Fatalf("retry after canceled initial authentication: %v", err)
	}
	authenticatedOpenstackTestConfig(t, retry)
	if got := keystone.tokenRequests.Load(); got != 2 {
		t.Errorf("token requests = %d, want a fresh authentication after cancellation (2)", got)
	}
	if provider.Meta() != nil {
		t.Error("canceled configuration mutated the shared provider singleton's Meta")
	}
}

func TestNoForkClientCacheConfigureUsesAuthenticatedTokenExpiry(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	initial := time.Now().UTC().Truncate(time.Second)
	var clock atomic.Int64
	clock.Store(initial.UnixNano())
	now := func() time.Time { return time.Unix(0, clock.Load()).UTC() }
	keystone := newKeystoneTestServer(t, 10*time.Minute, nil)
	keystone.now = now
	cache := newNoForkClientCache(8)
	cache.now = now
	provider := osprovider.Provider()
	configuration := keystoneTestConfiguration(keystone)
	configuration["allow_reauth"] = "false"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	first := terraform.Setup{Configuration: configuration}
	if err := cache.configure(ctx, &first, provider); err != nil {
		t.Fatalf("initial configure: %v", err)
	}
	firstConfig := authenticatedOpenstackTestConfig(t, first)
	if firstConfig.AllowReauth || firstConfig.OsClient.ReauthFunc != nil {
		t.Fatal("test requires allow_reauth=false, without automatic token refresh")
	}
	beforeExpiry := terraform.Setup{Configuration: configuration}
	if err := cache.configure(ctx, &beforeExpiry, provider); err != nil {
		t.Fatalf("configure before token expiry: %v", err)
	}
	if beforeExpiry.Meta != first.Meta || keystone.tokenRequests.Load() != 1 {
		t.Fatal("unexpired authenticated token was not cached")
	}
	// This is past the actual ten-minute token lifetime, but well before the
	// original thirty-minute fallback for an unauthenticated Config.
	clock.Store(initial.Add(11 * time.Minute).UnixNano())
	afterExpiry := terraform.Setup{Configuration: configuration}
	if err := cache.configure(ctx, &afterExpiry, provider); err != nil {
		t.Fatalf("configure after actual token expiry: %v", err)
	}
	fresh := authenticatedOpenstackTestConfig(t, afterExpiry)
	if fresh == firstConfig || fresh.OsClient == firstConfig.OsClient || fresh.OsClient.Token() == firstConfig.OsClient.Token() {
		t.Error("expired non-reauthenticating client was reused instead of freshly authenticated")
	}
	if got := keystone.tokenRequests.Load(); got != 2 {
		t.Errorf("token requests after actual expiry = %d, want 2", got)
	}
	if provider.Meta() != nil {
		t.Error("expiry refresh mutated the shared provider singleton's Meta")
	}
}

func TestNoForkClientCacheConfigureIsolatesCredentials(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	keystone := newKeystoneTestServer(t, time.Hour, nil)
	cache := newNoForkClientCache(8)
	provider := osprovider.Provider()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	credentials := []keystoneTestCredentials{
		{password: "dummy-password", tenantID: "dummy-project"},
		{password: "dummy-rotated-password", tenantID: "dummy-project"},
		{password: "dummy-password", tenantID: "dummy-other-project"},
	}
	configs := make([]*osprovider.Config, len(credentials))
	for round := range 2 {
		for i, credential := range credentials {
			configuration := keystoneTestConfiguration(keystone)
			configuration["password"] = credential.password
			configuration["tenant_id"] = credential.tenantID
			ps := terraform.Setup{Configuration: configuration}
			if err := cache.configure(ctx, &ps, provider); err != nil {
				t.Fatalf("configure credential set %d, round %d: %v", i, round, err)
			}
			cfg := authenticatedOpenstackTestConfig(t, ps)
			if round == 0 {
				for _, previous := range configs[:i] {
					if cfg == previous || cfg.OsClient == previous.OsClient || cfg.OsClient.Token() == previous.OsClient.Token() {
						t.Errorf("credential set %d shared a client or token with different credentials", i)
					}
				}
				configs[i] = cfg
			} else if cfg != configs[i] {
				t.Errorf("credential set %d did not reuse its own cached Config", i)
			}
		}
	}
	if got := keystone.tokenRequests.Load(); got != int32(len(credentials)) {
		t.Errorf("token requests = %d, want one per distinct credential set (%d)", got, len(credentials))
	}
	keystone.mu.Lock()
	defer keystone.mu.Unlock()
	if len(keystone.credentials) != len(credentials) {
		t.Fatalf("Keystone received %d credential sets, want %d", len(keystone.credentials), len(credentials))
	}
	for i, credential := range credentials {
		if keystone.credentials[i] != credential {
			t.Errorf("Keystone request %d did not use its configured password and project scope", i)
		}
	}
	if provider.Meta() != nil {
		t.Error("different credentials mutated the shared provider singleton's Meta")
	}
}

func TestNoForkClientCacheConfigureBypassesSwauth(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	keystone := newKeystoneTestServer(t, time.Hour, nil)
	cache := newNoForkClientCache(8)
	provider := osprovider.Provider()
	var previous any
	for range 2 {
		configuration := keystoneTestConfiguration(keystone)
		configuration["swauth"] = "true"
		ps := terraform.Setup{Configuration: configuration}
		if err := cache.configure(context.Background(), &ps, provider); err != nil {
			t.Fatalf("configure swauth: %v", err)
		}
		cfg, ok := ps.Meta.(*osprovider.Config)
		if !ok || cfg == nil || !cfg.Swauth {
			t.Fatalf("expected an actual swauth provider Config, got %T", ps.Meta)
		}
		if ps.Meta == previous {
			t.Error("swauth Config was cached")
		}
		previous = ps.Meta
	}
	if got := keystone.tokenRequests.Load(); got != 0 {
		t.Errorf("swauth configuration requested %d Keystone tokens, want 0", got)
	}
	if provider.Meta() != nil {
		t.Error("swauth configuration mutated the shared provider singleton's Meta")
	}
}

func TestNoForkClientCacheConfigureBypassesUnknownMeta(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	var calls int
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"insecure": {Type: schema.TypeBool, Optional: true},
		},
		ConfigureContextFunc: func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
			calls++
			return &struct{ call int }{call: calls}, nil
		},
	}
	cache := newNoForkClientCache(8)
	var previous any
	for range 2 {
		ps := terraform.Setup{Configuration: map[string]any{"insecure": "false"}}
		if err := cache.configure(context.Background(), &ps, provider); err != nil {
			t.Fatalf("configure unknown Meta: %v", err)
		}
		if ps.Meta == nil || ps.Meta == previous {
			t.Error("unknown provider Meta was cached or not configured")
		}
		previous = ps.Meta
	}
	if calls != 2 {
		t.Errorf("provider configure calls = %d, want 2 for uncacheable Meta", calls)
	}
	if provider.Meta() != nil {
		t.Error("unknown Meta configuration mutated the shared provider singleton's Meta")
	}
}

func TestTerraformSetupBuilderReusesAuthenticatedClient(t *testing.T) {
	clearOpenstackTestEnvironment(t)
	for _, flavor := range []string{"Cluster", "Namespaced", "NamespacedClusterProviderConfig"} {
		t.Run(flavor, func(t *testing.T) {
			keystone := newKeystoneTestServer(t, time.Hour, nil)
			crClient, managed := openstackTestKubernetesObjects(t, flavor, keystoneTestConfiguration(keystone))
			provider := osprovider.Provider()
			setup := TerraformSetupBuilder(provider)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			var shared *osprovider.Config
			// Two different managed resources share their ProviderConfig, and a
			// subsequent reconciliation of the first resource must reuse it too.
			for _, mg := range []resource.Managed{managed[0], managed[1], managed[0]} {
				ps, err := setup(ctx, crClient, mg)
				if err != nil {
					t.Fatalf("Terraform setup for %s: %v", mg.GetName(), err)
				}
				cfg := authenticatedOpenstackTestConfig(t, ps)
				if shared == nil {
					shared = cfg
				} else if cfg != shared {
					t.Error("resources sharing a ProviderConfig did not share their authenticated Config")
				}
			}
			if got := keystone.tokenRequests.Load(); got != 1 {
				t.Errorf("Terraform setups requested %d tokens, want 1", got)
			}
			if provider.Meta() != nil {
				t.Error("Terraform setup mutated the shared provider singleton's Meta")
			}
			if flavor == "Cluster" {
				usages := &clusterv1beta1.ProviderConfigUsageList{}
				if err := crClient.List(ctx, usages); err != nil {
					t.Fatalf("list cluster ProviderConfig usages: %v", err)
				}
				if len(usages.Items) != 2 {
					t.Errorf("cluster ProviderConfig usages = %d, want 2", len(usages.Items))
				}
			} else {
				usages := &namespacedv1beta1.ProviderConfigUsageList{}
				if err := crClient.List(ctx, usages, client.InNamespace("dummy-resources")); err != nil {
					t.Fatalf("list namespaced ProviderConfig usages: %v", err)
				}
				if len(usages.Items) != 2 {
					t.Errorf("namespaced ProviderConfig usages = %d, want 2", len(usages.Items))
				}
			}
		})
	}
}

func openstackTestKubernetesObjects(t *testing.T, flavor string, configuration map[string]any) (client.Client, []resource.Managed) {
	t.Helper()
	scheme := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{
		corev1.AddToScheme,
		clusterv1beta1.SchemeBuilder.AddToScheme,
		namespacedv1beta1.SchemeBuilder.AddToScheme,
		clustercomputev1alpha1.AddToScheme,
		namespacedcomputev1alpha1.AddToScheme,
	} {
		if err := add(scheme); err != nil {
			t.Fatalf("register test Kubernetes types: %v", err)
		}
	}
	data, err := json.Marshal(configuration)
	if err != nil {
		t.Fatalf("marshal dummy provider credentials: %v", err)
	}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: "dummy-credentials", Namespace: "dummy-system"},
		Data:       map[string][]byte{"credentials": data},
	}
	selectors := xpv1.CommonCredentialSelectors{SecretRef: &xpv1.SecretKeySelector{
		SecretReference: xpv1.SecretReference{Name: secret.Name, Namespace: secret.Namespace}, Key: "credentials",
	}}
	objects := []client.Object{secret}
	managed := make([]resource.Managed, 0, 2)
	if flavor == "Cluster" {
		objects = append(objects, &clusterv1beta1.ProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "dummy-provider"},
			Spec: clusterv1beta1.ProviderConfigSpec{Credentials: clusterv1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret, CommonCredentialSelectors: selectors,
			}},
		})
		for _, name := range []string{"dummy-resource-a", "dummy-resource-b"} {
			mg := &clustercomputev1alpha1.FlavorV2{
				TypeMeta:   metav1.TypeMeta{APIVersion: clustercomputev1alpha1.CRDGroupVersion.String(), Kind: "FlavorV2"},
				ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(name)},
			}
			mg.SetProviderConfigReference(&xpv1.Reference{Name: "dummy-provider"})
			managed = append(managed, mg)
			objects = append(objects, mg)
		}
	} else {
		kind := "ClusterProviderConfig"
		pcSpec := namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
			Source: xpv1.CredentialsSourceSecret, CommonCredentialSelectors: selectors,
		}}
		if flavor == "Namespaced" {
			kind = "ProviderConfig"
			secret.Namespace = "dummy-resources"
			// Namespaced ProviderConfig credentials must resolve in the managed
			// resource's namespace, not this deliberately incorrect selector.
			pcSpec.Credentials.SecretRef.Namespace = "dummy-ignored-namespace"
			objects = append(objects, &namespacedv1beta1.ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{Name: "dummy-provider", Namespace: "dummy-resources"}, Spec: pcSpec,
			})
		} else {
			objects = append(objects, &namespacedv1beta1.ClusterProviderConfig{
				ObjectMeta: metav1.ObjectMeta{Name: "dummy-provider"}, Spec: pcSpec,
			})
		}
		for _, name := range []string{"dummy-resource-a", "dummy-resource-b"} {
			mg := &namespacedcomputev1alpha1.FlavorV2{
				TypeMeta:   metav1.TypeMeta{APIVersion: namespacedcomputev1alpha1.CRDGroupVersion.String(), Kind: "FlavorV2"},
				ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: "dummy-resources", UID: types.UID(name)},
			}
			mg.SetProviderConfigReference(&xpv1.ProviderConfigReference{Name: "dummy-provider", Kind: kind})
			managed = append(managed, mg)
			objects = append(objects, mg)
		}
	}
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objects...).WithInterceptorFuncs(interceptor.Funcs{
		Get: func(ctx context.Context, c client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
			// The real API client ignores a namespace in a cluster-scoped GET;
			// the fake's object tracker does not know resource scopes.
			if _, ok := obj.(*namespacedv1beta1.ClusterProviderConfig); ok {
				key.Namespace = ""
			}
			return c.Get(ctx, key, obj, opts...)
		},
	}).Build(), managed
}
