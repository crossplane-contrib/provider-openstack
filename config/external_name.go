/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda

	"openstack_compute_instance_v2": config.IdentifierFromProvider,
	"openstack_compute_keypair_v2":  config.IdentifierFromProvider,
	"openstack_compute_quotaset_v2": config.IdentifierFromProvider,

	"openstack_networking_network_v2":          config.IdentifierFromProvider,
	"openstack_networking_subnet_v2":           config.IdentifierFromProvider,
	"openstack_networking_router_v2":           config.IdentifierFromProvider,
	"openstack_networking_router_interface_v2": config.IdentifierFromProvider,
	"openstack_networking_floatingip_v2":       config.IdentifierFromProvider,

	"openstack_dns_zone_v2":      config.IdentifierFromProvider,
	"openstack_dns_recordset_v2": config.IdentifierFromProvider,

	"openstack_containerinfra_cluster_v1":         config.IdentifierFromProvider,
	"openstack_containerinfra_clustertemplate_v1": config.IdentifierFromProvider,
	"openstack_containerinfra_nodegroup_v1":       config.IdentifierFromProvider,

	"openstack_identity_project_v3":         config.IdentifierFromProvider,
	"openstack_identity_role_assignment_v3": config.IdentifierFromProvider,
	"openstack_identity_user_v3":            config.IdentifierFromProvider,

	"openstack_blockstorage_quotaset_v3": config.IdentifierFromProvider,

	"openstack_lb_quota_v2": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
