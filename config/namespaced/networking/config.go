package networking

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_networking_subnet_v2", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			TerraformName: "openstack_networking_network_v2",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"allocation_pools"},
		}
	})
	p.AddResourceConfigurator("openstack_networking_floatingip_associate_v2", func(r *config.Resource) {
		r.References["floating_ip"] = config.Reference{
			TerraformName: "openstack_networking_floatingip_v2",
			Extractor:     "github.com/crossplane-contrib/provider-openstack/apis/namespaced/networking/v1alpha1.ExtractFloatingIPAddress()",
		}
	})
	p.AddResourceConfigurator("openstack_networking_router_interface_v2", func(r *config.Resource) {
		r.References["router_id"] = config.Reference{
			TerraformName: "openstack_networking_router_v2",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "openstack_networking_subnet_v2",
		}
		r.References["port_id"] = config.Reference{
			TerraformName: "openstack_networking_port_v2",
		}
	})
	p.AddResourceConfigurator("openstack_networking_router_v2", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"external_gateway"},
		}
	})
	p.AddResourceConfigurator("openstack_networking_quota_v2", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
		}
	})
}
