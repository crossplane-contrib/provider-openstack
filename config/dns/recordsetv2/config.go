package recordsetv2

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_dns_recordset_v2", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "ZoneV2",
		}
	})
}
