/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/martinnirtl/provider-openstack/config/compute/instancev2"
	"github.com/martinnirtl/provider-openstack/config/dns/recordsetv2"
	"github.com/martinnirtl/provider-openstack/config/dns/zonev2"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "openstack"
	modulePath     = "github.com/martinnirtl/provider-openstack"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		compute.Configure,
		recordsetv2.Configure,
		zonev2.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
