package config

import (
	"context"
	"strings"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/terraform-provider-openstack/terraform-provider-openstack/v3/openstack"

	blockstorageCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/blockstorage"
	blockstorageNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/blockstorage"
	computeCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/compute"
	computeNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/compute"
	containerinfraCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/containerinfra"
	containerinfraNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/containerinfra"
	dnsCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/dns"
	dnsNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/dns"
	identityCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/identity"
	identityNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/identity"
	lbCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/lb"
	lbNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/lb"
	networkingCluster "github.com/crossplane-contrib/provider-openstack/config/cluster/networking"
	networkingNamespaced "github.com/crossplane-contrib/provider-openstack/config/namespaced/networking"
)

const (
	resourcePrefix = "openstack"
	modulePath     = "github.com/crossplane-contrib/provider-openstack"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// oldSingletonListAPIs is a newline-delimited list of Terraform resource
// names with converted singleton list APIs with at least CRD API version
// containing the old singleton list API. This is to prevent the API
// conversion for the newly added resources whose CRD APIs will already
// use embedded objects instead of the singleton lists and thus, will
// not possess a CRD API version with the singleton list. Thus, for
// the newly added resources (resources added after the singleton lists
// have been converted), we do not need the CRD API conversion
// functions that convert between singleton lists and embedded objects,
// but we need only the Terraform conversion functions.
// This list is immutable and represents the set of resources with the
// already generated CRD API versions with now converted singleton lists.
// Because new resources should never have singleton lists in their
// generated APIs, there should be no need to add them to this list.
// However, bugs might result in exceptions in the future.
// Please see:
// https://github.com/crossplane-contrib/provider-upjet-azuread/pull/123
// for more context on singleton list to embedded object conversions.
var oldSingletonListAPIs string

// workaround for the TF Google v2.41.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	sdkProvider := openstack.Provider()

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("openstack.crossplane.io"),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithDefaultResourceOptions(
			resourceConfigurator(),
		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	bumpVersionsWithEmbeddedLists(pc)
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		blockstorageCluster.Configure,
		computeCluster.Configure,
		containerinfraCluster.Configure,
		dnsCluster.Configure,
		identityCluster.Configure,
		lbCluster.Configure,
		networkingCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

func GetProviderNamespaced(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	sdkProvider := openstack.Provider()

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("openstack.m.crossplane.io"),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithDefaultResourceOptions(
			resourceConfigurator(),
		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	bumpVersionsWithEmbeddedLists(pc)
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		blockstorageNamespaced.Configure,
		computeNamespaced.Configure,
		containerinfraNamespaced.Configure,
		dnsNamespaced.Configure,
		identityNamespaced.Configure,
		lbNamespaced.Configure,
		networkingNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}

func bumpVersionsWithEmbeddedLists(pc *ujconfig.Provider) {
	l := strings.Split(strings.TrimSpace(oldSingletonListAPIs), "\n")
	oldSLAPIs := make(map[string]struct{}, len(l))
	for _, n := range l {
		oldSLAPIs[n] = struct{}{}
	}

	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		if _, ok := oldSLAPIs[name]; ok {
			r.Version = "v1alpha1"
			r.PreviousVersions = []string{"v1alpha1"}
			// we would like to set the storage version to v1alpha1 to facilitate
			// downgrades.
			r.SetCRDStorageVersion("v1alpha1")
			// because the controller reconciles on the API version with the singleton list API,
			// no need for a Terraform conversion.
			r.ControllerReconcileVersion = "v1alpha1"
		} else {
			// the controller will be reconciling on the CRD API version
			// with the converted API (with embedded objects in place of
			// singleton lists), so we need the appropriate Terraform
			// converter in this case.
			r.TerraformConversions = []ujconfig.TerraformConversion{
				ujconfig.NewTFSingletonConversion(),
			}
		}
		pc.Resources[name] = r
	}
}
