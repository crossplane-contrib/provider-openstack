package containerinfra

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_containerinfra_cluster_v1", func(r *config.Resource) {
		// r.References["fixed_network"] = config.Reference{
		// 	Type: "NetworkV2",
		// }
		// r.References["fixed_subnet"] = config.Reference{
		// 	Type: "SubnetV2",
		// }
	})
}
