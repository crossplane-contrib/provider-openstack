package dns

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_dns_zone_v2", func(r *config.Resource) {
		// nothing here atm
	})
	p.AddResourceConfigurator("openstack_dns_recordset_v2", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			TerraformName: "openstack_dns_zone_v2",
		}
	})
}
