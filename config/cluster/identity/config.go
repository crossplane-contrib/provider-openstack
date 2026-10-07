package identity

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-openstack/internal/initializer/domain"
)

const domainIDExtractor = "github.com/crossplane-contrib/provider-openstack/apis/cluster/identity/v1alpha1.ExtractDomainID()"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_identity_project_v3", func(r *config.Resource) {
		r.References["domain_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
			Extractor:     domainIDExtractor,
		}
		r.InitializerFns = append(r.InitializerFns, domain.NewValidator)
	})
	p.AddResourceConfigurator("openstack_identity_user_v3", func(r *config.Resource) {
		r.References["domain_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
			Extractor:     domainIDExtractor,
		}
		r.InitializerFns = append(r.InitializerFns, domain.NewValidator)
	})
	p.AddResourceConfigurator("openstack_identity_group_v3", func(r *config.Resource) {
		r.References["domain_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
			Extractor:     domainIDExtractor,
		}
		r.InitializerFns = append(r.InitializerFns, domain.NewValidator)
	})
	p.AddResourceConfigurator("openstack_identity_role_v3", func(r *config.Resource) {
		r.References["domain_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
			Extractor:     domainIDExtractor,
		}
		r.InitializerFns = append(r.InitializerFns, domain.NewValidator)
	})
	p.AddResourceConfigurator("openstack_identity_role_assignment_v3", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
		}
		r.References["domain_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
			Extractor:     domainIDExtractor,
		}
		r.InitializerFns = append(r.InitializerFns, domain.NewValidator)
	})
	p.AddResourceConfigurator("openstack_identity_inherit_role_assignment_v3", func(r *config.Resource) {
		r.References["domain_id"] = config.Reference{
			TerraformName: "openstack_identity_project_v3",
			Extractor:     domainIDExtractor,
		}
		r.InitializerFns = append(r.InitializerFns, domain.NewValidator)
	})
}
