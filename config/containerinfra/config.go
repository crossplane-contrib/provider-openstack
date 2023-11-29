package containerinfra

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_containerinfra_cluster_v1", func(r *config.Resource) {
		r.References["fixed_network"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.NetworkV2",
		}
		r.References["fixed_subnet"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2",
		}
	})
}
