package blockstorage

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_blockstorage_quotaset_v3", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
		}
	})
}
