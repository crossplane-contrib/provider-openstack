/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-openstack/config/blockstorage"
	"github.com/crossplane-contrib/provider-openstack/config/compute"
	"github.com/crossplane-contrib/provider-openstack/config/containerinfra"
	"github.com/crossplane-contrib/provider-openstack/config/dns"
	"github.com/crossplane-contrib/provider-openstack/config/identity"
	"github.com/crossplane-contrib/provider-openstack/config/lb"
	"github.com/crossplane-contrib/provider-openstack/config/networking"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "openstack"
	modulePath     = "github.com/crossplane-contrib/provider-openstack"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("openstack.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		blockstorage.Configure,
		compute.Configure,
		containerinfra.Configure,
		dns.Configure,
		identity.Configure,
		lb.Configure,
		networking.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
