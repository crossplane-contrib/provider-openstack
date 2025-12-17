/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
Copyright 2025 Yannick Schlosser, Jan Dittrich
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	qosassociationv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/qosassociationv3"
	qosv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/qosv3"
	quotasetv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/quotasetv3"
	volumeattachv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/volumeattachv3"
	volumetypeaccessv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/volumetypeaccessv3"
	volumetypev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/volumetypev3"
	volumev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/blockstorage/volumev3"
	aggregatev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/aggregatev2"
	flavoraccessv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/flavoraccessv2"
	flavorv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/flavorv2"
	instancev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/instancev2"
	interfaceattachv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/interfaceattachv2"
	keypairv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/keypairv2"
	quotasetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/quotasetv2"
	servergroupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/servergroupv2"
	volumeattachv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/compute/volumeattachv2"
	clustertemplatev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/containerinfra/clustertemplatev1"
	clusterv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/containerinfra/clusterv1"
	nodegroupv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/containerinfra/nodegroupv1"
	configurationv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/db/configurationv1"
	databasev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/db/databasev1"
	instancev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/db/instancev1"
	userv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/db/userv1"
	recordsetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/dns/recordsetv2"
	transferacceptv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/dns/transferacceptv2"
	transferrequestv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/dns/transferrequestv2"
	zonev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/dns/zonev2"
	groupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/fw/groupv2"
	policyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/fw/policyv2"
	rulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/fw/rulev2"
	applicationcredentialv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/applicationcredentialv3"
	ec2credentialv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/ec2credentialv3"
	endpointv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/endpointv3"
	groupv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/groupv3"
	inheritroleassignmentv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/inheritroleassignmentv3"
	projectv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/projectv3"
	roleassignmentv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/roleassignmentv3"
	rolev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/rolev3"
	servicev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/servicev3"
	usermembershipv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/usermembershipv3"
	userv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/identity/userv3"
	imageaccessacceptv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/images/imageaccessacceptv2"
	imageaccessv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/images/imageaccessv2"
	imagev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/images/imagev2"
	containerv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/keymanager/containerv1"
	orderv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/keymanager/orderv1"
	secretv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/keymanager/secretv1"
	l7policyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/l7policyv2"
	l7rulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/l7rulev2"
	listenerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/listenerv2"
	loadbalancerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/loadbalancerv2"
	membersv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/membersv2"
	memberv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/memberv2"
	monitorv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/monitorv2"
	poolv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/poolv2"
	quotav2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/lb/quotav2"
	addressscopev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/addressscopev2"
	floatingipassociatev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/floatingipassociatev2"
	floatingipv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/floatingipv2"
	networkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/networkv2"
	portforwardingv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/portforwardingv2"
	portsecgroupassociatev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/portsecgroupassociatev2"
	portv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/portv2"
	qosbandwidthlimitrulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/qosbandwidthlimitrulev2"
	qosdscpmarkingrulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/qosdscpmarkingrulev2"
	qosminimumbandwidthrulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/qosminimumbandwidthrulev2"
	qospolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/qospolicyv2"
	quotav2networking "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/quotav2"
	rbacpolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/rbacpolicyv2"
	routerinterfacev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/routerinterfacev2"
	routerroutev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/routerroutev2"
	routerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/routerv2"
	secgrouprulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/secgrouprulev2"
	secgroupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/secgroupv2"
	subnetpoolv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/subnetpoolv2"
	subnetroutev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/subnetroutev2"
	subnetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/subnetv2"
	trunkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/networking/trunkv2"
	containerv1objectstorage "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/objectstorage/containerv1"
	objectv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/objectstorage/objectv1"
	tempurlv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/objectstorage/tempurlv1"
	stackv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/orchestration/stackv1"
	providerconfig "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/providerconfig"
	securityservicev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/sharedfilesystem/securityservicev2"
	shareaccessv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/sharedfilesystem/shareaccessv2"
	sharenetworkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/sharedfilesystem/sharenetworkv2"
	sharev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/sharedfilesystem/sharev2"
	endpointgroupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/vpnaas/endpointgroupv2"
	ikepolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/vpnaas/ikepolicyv2"
	ipsecpolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/vpnaas/ipsecpolicyv2"
	servicev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/vpnaas/servicev2"
	siteconnectionv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster/vpnaas/siteconnectionv2"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		qosassociationv3.Setup,
		qosv3.Setup,
		quotasetv3.Setup,
		volumeattachv3.Setup,
		volumetypeaccessv3.Setup,
		volumetypev3.Setup,
		volumev3.Setup,
		aggregatev2.Setup,
		flavoraccessv2.Setup,
		flavorv2.Setup,
		instancev2.Setup,
		interfaceattachv2.Setup,
		keypairv2.Setup,
		quotasetv2.Setup,
		servergroupv2.Setup,
		volumeattachv2.Setup,
		clustertemplatev1.Setup,
		clusterv1.Setup,
		nodegroupv1.Setup,
		configurationv1.Setup,
		databasev1.Setup,
		instancev1.Setup,
		userv1.Setup,
		recordsetv2.Setup,
		transferacceptv2.Setup,
		transferrequestv2.Setup,
		zonev2.Setup,
		groupv2.Setup,
		policyv2.Setup,
		rulev2.Setup,
		applicationcredentialv3.Setup,
		ec2credentialv3.Setup,
		endpointv3.Setup,
		groupv3.Setup,
		inheritroleassignmentv3.Setup,
		projectv3.Setup,
		roleassignmentv3.Setup,
		rolev3.Setup,
		servicev3.Setup,
		usermembershipv3.Setup,
		userv3.Setup,
		imageaccessacceptv2.Setup,
		imageaccessv2.Setup,
		imagev2.Setup,
		containerv1.Setup,
		orderv1.Setup,
		secretv1.Setup,
		l7policyv2.Setup,
		l7rulev2.Setup,
		listenerv2.Setup,
		loadbalancerv2.Setup,
		membersv2.Setup,
		memberv2.Setup,
		monitorv2.Setup,
		poolv2.Setup,
		quotav2.Setup,
		addressscopev2.Setup,
		floatingipassociatev2.Setup,
		floatingipv2.Setup,
		networkv2.Setup,
		portforwardingv2.Setup,
		portsecgroupassociatev2.Setup,
		portv2.Setup,
		qosbandwidthlimitrulev2.Setup,
		qosdscpmarkingrulev2.Setup,
		qosminimumbandwidthrulev2.Setup,
		qospolicyv2.Setup,
		quotav2networking.Setup,
		rbacpolicyv2.Setup,
		routerinterfacev2.Setup,
		routerroutev2.Setup,
		routerv2.Setup,
		secgrouprulev2.Setup,
		secgroupv2.Setup,
		subnetpoolv2.Setup,
		subnetroutev2.Setup,
		subnetv2.Setup,
		trunkv2.Setup,
		containerv1objectstorage.Setup,
		objectv1.Setup,
		tempurlv1.Setup,
		stackv1.Setup,
		providerconfig.Setup,
		securityservicev2.Setup,
		shareaccessv2.Setup,
		sharenetworkv2.Setup,
		sharev2.Setup,
		endpointgroupv2.Setup,
		ikepolicyv2.Setup,
		ipsecpolicyv2.Setup,
		servicev2.Setup,
		siteconnectionv2.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		qosassociationv3.SetupGated,
		qosv3.SetupGated,
		quotasetv3.SetupGated,
		volumeattachv3.SetupGated,
		volumetypeaccessv3.SetupGated,
		volumetypev3.SetupGated,
		volumev3.SetupGated,
		aggregatev2.SetupGated,
		flavoraccessv2.SetupGated,
		flavorv2.SetupGated,
		instancev2.SetupGated,
		interfaceattachv2.SetupGated,
		keypairv2.SetupGated,
		quotasetv2.SetupGated,
		servergroupv2.SetupGated,
		volumeattachv2.SetupGated,
		clustertemplatev1.SetupGated,
		clusterv1.SetupGated,
		nodegroupv1.SetupGated,
		configurationv1.SetupGated,
		databasev1.SetupGated,
		instancev1.SetupGated,
		userv1.SetupGated,
		recordsetv2.SetupGated,
		transferacceptv2.SetupGated,
		transferrequestv2.SetupGated,
		zonev2.SetupGated,
		groupv2.SetupGated,
		policyv2.SetupGated,
		rulev2.SetupGated,
		applicationcredentialv3.SetupGated,
		ec2credentialv3.SetupGated,
		endpointv3.SetupGated,
		groupv3.SetupGated,
		inheritroleassignmentv3.SetupGated,
		projectv3.SetupGated,
		roleassignmentv3.SetupGated,
		rolev3.SetupGated,
		servicev3.SetupGated,
		usermembershipv3.SetupGated,
		userv3.SetupGated,
		imageaccessacceptv2.SetupGated,
		imageaccessv2.SetupGated,
		imagev2.SetupGated,
		containerv1.SetupGated,
		orderv1.SetupGated,
		secretv1.SetupGated,
		l7policyv2.SetupGated,
		l7rulev2.SetupGated,
		listenerv2.SetupGated,
		loadbalancerv2.SetupGated,
		membersv2.SetupGated,
		memberv2.SetupGated,
		monitorv2.SetupGated,
		poolv2.SetupGated,
		quotav2.SetupGated,
		addressscopev2.SetupGated,
		floatingipassociatev2.SetupGated,
		floatingipv2.SetupGated,
		networkv2.SetupGated,
		portforwardingv2.SetupGated,
		portsecgroupassociatev2.SetupGated,
		portv2.SetupGated,
		qosbandwidthlimitrulev2.SetupGated,
		qosdscpmarkingrulev2.SetupGated,
		qosminimumbandwidthrulev2.SetupGated,
		qospolicyv2.SetupGated,
		quotav2networking.SetupGated,
		rbacpolicyv2.SetupGated,
		routerinterfacev2.SetupGated,
		routerroutev2.SetupGated,
		routerv2.SetupGated,
		secgrouprulev2.SetupGated,
		secgroupv2.SetupGated,
		subnetpoolv2.SetupGated,
		subnetroutev2.SetupGated,
		subnetv2.SetupGated,
		trunkv2.SetupGated,
		containerv1objectstorage.SetupGated,
		objectv1.SetupGated,
		tempurlv1.SetupGated,
		stackv1.SetupGated,
		providerconfig.SetupGated,
		securityservicev2.SetupGated,
		shareaccessv2.SetupGated,
		sharenetworkv2.SetupGated,
		sharev2.SetupGated,
		endpointgroupv2.SetupGated,
		ikepolicyv2.SetupGated,
		ipsecpolicyv2.SetupGated,
		servicev2.SetupGated,
		siteconnectionv2.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
