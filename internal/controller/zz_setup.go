// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	qosassociationv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/qosassociationv3"
	qosv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/qosv3"
	quotasetv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/quotasetv3"
	volumeattachv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/volumeattachv3"
	volumetypeaccessv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/volumetypeaccessv3"
	volumetypev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/volumetypev3"
	volumev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/blockstorage/volumev3"
	aggregatev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/aggregatev2"
	flavoraccessv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/flavoraccessv2"
	flavorv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/flavorv2"
	floatingipassociatev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/floatingipassociatev2"
	floatingipv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/floatingipv2"
	instancev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/instancev2"
	interfaceattachv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/interfaceattachv2"
	keypairv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/keypairv2"
	quotasetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/quotasetv2"
	secgroupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/secgroupv2"
	servergroupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/servergroupv2"
	volumeattachv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/volumeattachv2"
	clustertemplatev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/containerinfra/clustertemplatev1"
	clusterv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/containerinfra/clusterv1"
	nodegroupv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/containerinfra/nodegroupv1"
	configurationv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/db/configurationv1"
	databasev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/db/databasev1"
	instancev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/db/instancev1"
	userv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/db/userv1"
	recordsetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/dns/recordsetv2"
	transferacceptv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/dns/transferacceptv2"
	transferrequestv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/dns/transferrequestv2"
	zonev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/dns/zonev2"
	groupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/fw/groupv2"
	policyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/fw/policyv2"
	rulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/fw/rulev2"
	applicationcredentialv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/applicationcredentialv3"
	ec2credentialv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/ec2credentialv3"
	endpointv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/endpointv3"
	groupv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/groupv3"
	inheritroleassignmentv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/inheritroleassignmentv3"
	projectv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/projectv3"
	roleassignmentv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/roleassignmentv3"
	rolev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/rolev3"
	servicev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/servicev3"
	usermembershipv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/usermembershipv3"
	userv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/userv3"
	imageaccessacceptv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/images/imageaccessacceptv2"
	imageaccessv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/images/imageaccessv2"
	imagev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/images/imagev2"
	containerv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/keymanager/containerv1"
	orderv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/keymanager/orderv1"
	secretv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/keymanager/secretv1"
	l7policyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/l7policyv2"
	l7rulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/l7rulev2"
	listenerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/listenerv2"
	loadbalancerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/loadbalancerv2"
	membersv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/membersv2"
	memberv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/memberv2"
	monitorv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/monitorv2"
	poolv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/poolv2"
	quotav2 "github.com/crossplane-contrib/provider-openstack/internal/controller/lb/quotav2"
	addressscopev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/addressscopev2"
	floatingipassociatev2networking "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/floatingipassociatev2"
	floatingipv2networking "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/floatingipv2"
	networkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/networkv2"
	portforwardingv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/portforwardingv2"
	portsecgroupassociatev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/portsecgroupassociatev2"
	portv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/portv2"
	qosbandwidthlimitrulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/qosbandwidthlimitrulev2"
	qosdscpmarkingrulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/qosdscpmarkingrulev2"
	qosminimumbandwidthrulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/qosminimumbandwidthrulev2"
	qospolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/qospolicyv2"
	quotav2networking "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/quotav2"
	rbacpolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/rbacpolicyv2"
	routerinterfacev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerinterfacev2"
	routerroutev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerroutev2"
	routerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerv2"
	secgrouprulev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/secgrouprulev2"
	secgroupv2networking "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/secgroupv2"
	subnetpoolv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/subnetpoolv2"
	subnetroutev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/subnetroutev2"
	subnetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/subnetv2"
	trunkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/trunkv2"
	containerv1objectstorage "github.com/crossplane-contrib/provider-openstack/internal/controller/objectstorage/containerv1"
	objectv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/objectstorage/objectv1"
	tempurlv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/objectstorage/tempurlv1"
	stackv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/orchestration/stackv1"
	providerconfig "github.com/crossplane-contrib/provider-openstack/internal/controller/providerconfig"
	securityservicev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/sharedfilesystem/securityservicev2"
	shareaccessv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/sharedfilesystem/shareaccessv2"
	sharenetworkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/sharedfilesystem/sharenetworkv2"
	sharev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/sharedfilesystem/sharev2"
	endpointgroupv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/vpnaas/endpointgroupv2"
	ikepolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/vpnaas/ikepolicyv2"
	ipsecpolicyv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/vpnaas/ipsecpolicyv2"
	servicev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/vpnaas/servicev2"
	siteconnectionv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/vpnaas/siteconnectionv2"
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
		floatingipassociatev2.Setup,
		floatingipv2.Setup,
		instancev2.Setup,
		interfaceattachv2.Setup,
		keypairv2.Setup,
		quotasetv2.Setup,
		secgroupv2.Setup,
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
		floatingipassociatev2networking.Setup,
		floatingipv2networking.Setup,
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
		secgroupv2networking.Setup,
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
