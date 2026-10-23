package clients

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	osprovider "github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/crossplane-contrib/provider-openstack/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/crossplane-contrib/provider-openstack/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal openstack credentials as JSON"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
//
// The returned SetupFn caches configured provider clients per provider
// configuration (see noForkClientCache), so that managed resources sharing a
// ProviderConfig reuse one authenticated OpenStack client across
// reconciliations instead of re-authenticating against the identity service
// on every reconciliation.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn { //nolint:gocyclo
	// The cache is scoped to this SetupFn, which is instantiated once per
	// provider flavor (cluster-scoped and namespaced), so entries of
	// different flavors never mix.
	cache := newNoForkClientCache(defaultNoForkClientCacheSize)

	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		// Keep in sync with configuration options of Terraform OpenStack provider:
		// https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs
		// Exceptions:
		// "cloud" is not used, because we dont support giving a clouds.yaml directly to the provider.
		credFields := []string{"auth_url", "region", "user_name", "user_id", "application_credential_id", "application_credential_name", "application_credential_secret",
			"tenant_id", "tenant_name", "password", "token", "user_domain_name", "user_domain_id", "project_domain_name", "project_domain_id", "domain_id", "domain_name",
			"default_domain", "system_scope", "insecure", "cacert_file", "cert", "key", "endpoint_type", "endpoint_overrides", "swauth", "use_octavia", "disable_no_cache_header",
			"delayed_auth", "allow_reauth", "max_retries", "enable_logging"}

		ps.Configuration = map[string]any{}
		// Provide default values for specific fields. They are the same as in the upstream provider.
		ps.Configuration["insecure"] = "false"

		// ensures only the provided fields are set in the config
		for _, credField := range credFields {
			if v, ok := creds[credField]; ok {
				ps.Configuration[credField] = v
			}
		}

		err = cache.configure(ctx, &ps, tfProvider)
		return ps, errors.Wrap(err, "failed to configure the no-fork OpenStack client")
	}
}

func (c *noForkClientCache) configure(ctx context.Context, ps *terraform.Setup, tfProvider *schema.Provider) error {
	configuration := ps.Configuration
	// Swift's non-delayed path writes directly to the ProviderClient token.
	// Do not share it even through singleflight while configuration is in
	// progress. Read its schema default too, to respect OS_SWAUTH.
	swauth, err := swiftAuthenticationEnabled(configuration, tfProvider)
	if err != nil {
		return errors.Wrap(err, "failed to parse 'swauth' configuration")
	}
	if swauth {
		return configureNoForkOpenstackClient(ctx, ps, tfProvider)
	}

	meta, err := c.getOrCreate(ctx, configuration, func() (any, error) {
		// Keep this setup local to the creator. A caller canceled while waiting
		// must not race with a later write from the initialization goroutine.
		fresh := terraform.Setup{Configuration: configuration}
		if err := configureNoForkOpenstackClient(ctx, &fresh, tfProvider); err != nil {
			return nil, err
		}
		return fresh.Meta, nil
	})
	if err != nil {
		return err
	}
	ps.Meta = meta
	return nil
}

func swiftAuthenticationEnabled(configuration map[string]any, tfProvider *schema.Provider) (bool, error) {
	value, ok := configuration["swauth"]
	if !ok && tfProvider.Schema["swauth"] != nil {
		var err error
		value, err = tfProvider.Schema["swauth"].DefaultValue()
		if err != nil {
			return false, err
		}
	}
	switch v := value.(type) {
	case string:
		return strconv.ParseBool(v)
	case bool:
		return v, nil
	default:
		return false, nil
	}
}

func configureNoForkOpenstackClient(ctx context.Context, ps *terraform.Setup, tfProvider *schema.Provider) error {
	// Configure a copy of the Terraform provider so that its Meta value
	// (holding the authenticated OpenStack client) can be cached and reused
	// across reconciliations, without mutating the shared provider singleton
	// whose Meta must stay unset for future configurations with different
	// credentials.
	p := *tfProvider

	// Populate the config, as well as the raw config for compatibility.
	config := tfsdk.NewResourceConfigRaw(ps.Configuration)

	// The insecure flag is checked using getOkExists, which uses the CtyType instead of config/rawconfig.
	// See https://github.com/terraform-provider-openstack/terraform-provider-openstack/blob/v3.0.0/openstack/provider.go#L577
	// To ensure compatibility with the upstream provider, we need to set the insecure flag in the cty.Value format.
	insecure, err := strconv.ParseBool(ps.Configuration["insecure"].(string))
	if err != nil {
		return errors.Wrap(err, "failed to parse 'insecure' configuration")
	}

	config.CtyValue = cty.ObjectVal(map[string]cty.Value{
		"insecure": cty.BoolVal(insecure),
	})

	diag := p.Configure(ctx, config)
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	meta := p.Meta()
	if cfg, ok := meta.(*osprovider.Config); ok && cfg != nil && !cfg.Swauth {
		// Delayed authentication normally happens during the first resource
		// operation. Finish it before publishing a shared client: upstream
		// retains initial authentication errors on Config permanently, and
		// token expiry is unavailable until authentication succeeds.
		if err := cfg.Authenticate(ctx); err != nil {
			return errors.Wrap(err, "failed to authenticate the OpenStack client")
		}
	}
	// openstack.NewClient already initializes Gophercloud's token/reauth locks.
	ps.Meta = meta
	return nil
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
