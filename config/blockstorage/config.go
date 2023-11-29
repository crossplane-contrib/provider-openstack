package blockstorage

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_blockstorage_quotaset_v3", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1.ProjectV3",
		}
	})
}
