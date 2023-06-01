package networking

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_networking_subnet_v2", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: "NetworkV2",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"allocation_pools"},
		}
	})
	p.AddResourceConfigurator("openstack_networking_router_interface_v2", func(r *config.Resource) {
		r.References["router_id"] = config.Reference{
			Type: "RouterV2",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "SubnetV2",
		}
	})
	p.AddResourceConfigurator("openstack_networking_router_v2", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"external_gateway"},
		}
	})
}
