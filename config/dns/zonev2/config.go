package zonev2

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_dns_zone_v2", func(r *config.Resource) {
		// nothing here atm
	})
}
