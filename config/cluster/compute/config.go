package compute

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_compute_instance_v2", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		// r.Name = "openstack_compute_instance_v2"
		// r.ExternalName = config.ExternalName{}
		// r.ShortGroup = "openstack"
		// r.Version = "v1alpha22"
		// Terraform state always has both flavor_id and flavor_name populated.
		// Late initialization of the counterpart field can introduce conflicting
		// desired state and cause resize loops during reconciliation.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"flavor_id", "flavor_name"},
		}
		r.References["key_pair"] = config.Reference{
			TerraformName: "openstack_compute_keypair_v2",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName: "openstack_networking_secgroup_v2",
		}
	})
	p.AddResourceConfigurator("openstack_compute_quotaset_v2", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
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
