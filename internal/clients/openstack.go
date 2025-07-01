package clients

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-openstack/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal openstack credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := xpresource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := xpresource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
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
		return ps, errors.Wrap(configureNoForkOpenstackClient(ctx, &ps, *tfProvider), "failed to configure the no-fork OpenStack client")
	}
}

func configureNoForkOpenstackClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
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

	diag := p.Configure(context.WithoutCancel(ctx), config)
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}
