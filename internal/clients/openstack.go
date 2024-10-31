/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

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
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
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

		// ensures only the provided fields are set in the config
		for _, credField := range credFields {
			if v, ok := creds[credField]; ok {
				ps.Configuration[credField] = v
			}
		}
		return ps, nil
	}
}
