package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var terraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	// Block Storage / Cinder
	"openstack_blockstorage_qos_association_v3":    config.IdentifierFromProvider,
	"openstack_blockstorage_qos_v3":                config.IdentifierFromProvider,
	"openstack_blockstorage_quotaset_v3":           config.IdentifierFromProvider,
	"openstack_blockstorage_volume_attach_v3":      config.IdentifierFromProvider,
	"openstack_blockstorage_volume_type_access_v3": config.IdentifierFromProvider,
	"openstack_blockstorage_volume_type_v3":        config.IdentifierFromProvider,
	"openstack_blockstorage_volume_v3":             config.IdentifierFromProvider,

	// Compute / Nova
	"openstack_compute_aggregate_v2":            config.IdentifierFromProvider,
	"openstack_compute_flavor_access_v2":        config.IdentifierFromProvider,
	"openstack_compute_flavor_v2":               config.IdentifierFromProvider,
	"openstack_compute_instance_v2":             config.IdentifierFromProvider,
	"openstack_compute_interface_attach_v2":     config.IdentifierFromProvider,
	"openstack_compute_keypair_v2":              config.IdentifierFromProvider,
	"openstack_compute_quotaset_v2":             config.IdentifierFromProvider,
	"openstack_compute_servergroup_v2":          config.IdentifierFromProvider,
	"openstack_compute_volume_attach_v2":        config.IdentifierFromProvider,

	// Container Infra / Magnum
	"openstack_containerinfra_cluster_v1":         config.IdentifierFromProvider,
	"openstack_containerinfra_clustertemplate_v1": config.IdentifierFromProvider,
	"openstack_containerinfra_nodegroup_v1":       config.IdentifierFromProvider,

	// DNS / Designate
	"openstack_dns_recordset_v2":        config.IdentifierFromProvider,
	"openstack_dns_transfer_accept_v2":  config.IdentifierFromProvider,
	"openstack_dns_transfer_request_v2": config.IdentifierFromProvider,
	"openstack_dns_zone_v2":             config.IdentifierFromProvider,

	// Databases / Trove
	"openstack_db_configuration_v1": config.IdentifierFromProvider,
	"openstack_db_database_v1":      config.IdentifierFromProvider,
	"openstack_db_instance_v1":      config.IdentifierFromProvider,
	"openstack_db_user_v1":          config.IdentifierFromProvider,

	// FWaaS / Neutron
	"openstack_fw_group_v2":  config.IdentifierFromProvider,
	"openstack_fw_policy_v2": config.IdentifierFromProvider,
	"openstack_fw_rule_v2":   config.IdentifierFromProvider,

	// Identity / Keystone
	"openstack_identity_application_credential_v3":  config.IdentifierFromProvider,
	"openstack_identity_ec2_credential_v3":          config.IdentifierFromProvider,
	"openstack_identity_endpoint_v3":                config.IdentifierFromProvider,
	"openstack_identity_group_v3":                   config.IdentifierFromProvider,
	"openstack_identity_inherit_role_assignment_v3": config.IdentifierFromProvider,
	"openstack_identity_project_v3":                 config.IdentifierFromProvider,
	"openstack_identity_role_assignment_v3":         config.IdentifierFromProvider,
	"openstack_identity_role_v3":                    config.IdentifierFromProvider,
	"openstack_identity_service_v3":                 config.IdentifierFromProvider,
	"openstack_identity_user_membership_v3":         config.IdentifierFromProvider,
	"openstack_identity_user_v3":                    config.IdentifierFromProvider,

	// Images / Glance
	"openstack_images_image_access_accept_v2": config.IdentifierFromProvider,
	"openstack_images_image_access_v2":        config.IdentifierFromProvider,
	"openstack_images_image_v2":               config.IdentifierFromProvider,

	// Key Manager / Barbican
	"openstack_keymanager_container_v1": config.IdentifierFromProvider,
	"openstack_keymanager_order_v1":     config.IdentifierFromProvider,
	"openstack_keymanager_secret_v1":    config.IdentifierFromProvider,

	// Load Balancing as a Service / Octavia
	"openstack_lb_l7policy_v2":     config.IdentifierFromProvider,
	"openstack_lb_l7rule_v2":       config.IdentifierFromProvider,
	"openstack_lb_listener_v2":     config.IdentifierFromProvider,
	"openstack_lb_loadbalancer_v2": config.IdentifierFromProvider,
	"openstack_lb_member_v2":       config.IdentifierFromProvider,
	"openstack_lb_members_v2":      config.IdentifierFromProvider,
	"openstack_lb_monitor_v2":      config.IdentifierFromProvider,
	"openstack_lb_pool_v2":         config.IdentifierFromProvider,
	"openstack_lb_quota_v2":        config.IdentifierFromProvider,

	// Networking / Neutron
	"openstack_networking_addressscope_v2":               config.IdentifierFromProvider,
	"openstack_networking_floatingip_associate_v2":       config.IdentifierFromProvider,
	"openstack_networking_floatingip_v2":                 config.IdentifierFromProvider,
	"openstack_networking_network_v2":                    config.IdentifierFromProvider,
	"openstack_networking_port_secgroup_associate_v2":    config.IdentifierFromProvider,
	"openstack_networking_port_v2":                       config.IdentifierFromProvider,
	"openstack_networking_portforwarding_v2":             config.IdentifierFromProvider,
	"openstack_networking_qos_bandwidth_limit_rule_v2":   config.IdentifierFromProvider,
	"openstack_networking_qos_dscp_marking_rule_v2":      config.IdentifierFromProvider,
	"openstack_networking_qos_minimum_bandwidth_rule_v2": config.IdentifierFromProvider,
	"openstack_networking_qos_policy_v2":                 config.IdentifierFromProvider,
	"openstack_networking_quota_v2":                      config.IdentifierFromProvider,
	"openstack_networking_rbac_policy_v2":                config.IdentifierFromProvider,
	"openstack_networking_router_interface_v2":           config.IdentifierFromProvider,
	"openstack_networking_router_route_v2":               config.IdentifierFromProvider,
	"openstack_networking_router_v2":                     config.IdentifierFromProvider,
	"openstack_networking_secgroup_rule_v2":              config.IdentifierFromProvider,
	"openstack_networking_secgroup_v2":                   config.IdentifierFromProvider,
	"openstack_networking_subnet_route_v2":               config.IdentifierFromProvider,
	"openstack_networking_subnet_v2":                     config.IdentifierFromProvider,
	"openstack_networking_subnetpool_v2":                 config.IdentifierFromProvider,
	"openstack_networking_trunk_v2":                      config.IdentifierFromProvider,

	// Object Storage / Swift
	"openstack_objectstorage_container_v1": config.IdentifierFromProvider,
	"openstack_objectstorage_object_v1":    config.IdentifierFromProvider,
	"openstack_objectstorage_tempurl_v1":   config.IdentifierFromProvider,

	// Orchestration / Heat
	"openstack_orchestration_stack_v1": config.IdentifierFromProvider,

	// Shared Filesystem / Manila
	"openstack_sharedfilesystem_securityservice_v2": config.IdentifierFromProvider,
	"openstack_sharedfilesystem_share_access_v2":    config.IdentifierFromProvider,
	"openstack_sharedfilesystem_share_v2":           config.IdentifierFromProvider,
	"openstack_sharedfilesystem_sharenetwork_v2":    config.IdentifierFromProvider,

	// VPNaaS / Neutron
	"openstack_vpnaas_endpoint_group_v2":  config.IdentifierFromProvider,
	"openstack_vpnaas_ike_policy_v2":      config.IdentifierFromProvider,
	"openstack_vpnaas_ipsec_policy_v2":    config.IdentifierFromProvider,
	"openstack_vpnaas_service_v2":         config.IdentifierFromProvider,
	"openstack_vpnaas_site_connection_v2": config.IdentifierFromProvider,
}

// cliReconciledExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the CLI-based
// architecture for this provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

// resourceConfigurator applies all external name configs
// listed in the table terraformPluginSDKExternalNameConfigs and
// cliReconciledExternalNameConfigs and sets the version
// of those resources to v1alpha1. For those resource in
// terraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func resourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := terraformPluginSDKExternalNameConfigs[r.Name]
		if !configured {
			e, configured = cliReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.Version = "v1alpha1"
		r.ExternalName = e
	}
}
