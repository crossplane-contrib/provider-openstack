package clusterv1

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_containerinfra_cluster_v1", func(r *config.Resource) {

	})
}
