package compute

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
	p.AddResourceConfigurator("openstack_compute_keypair_v2", func(r *config.Resource) {

		// example for connection secrets
		// r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
		// 	a := map[string][]byte{
		// 		"example secret": []byte("test"),
		// 	}
		// 	return a, nil
		// }
	})
}
