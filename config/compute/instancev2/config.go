package instancev2

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_compute_instance_v2", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		// r.Name = "openstack_compute_instance_v2"
		// r.ExternalName = config.ExternalName{}
		// r.ShortGroup = "openstack"
		// r.Version = "v1alpha22"
		r.References["key_pair"] = config.Reference{
			Type: "KeypairV2",
		}
	})
}
