/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	instancev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/instancev2"
	keypairv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/compute/keypairv2"
	clustertemplatev1 "github.com/crossplane-contrib/provider-openstack/internal/controller/containerinfra/clustertemplatev1"
	clusterv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/containerinfra/clusterv1"
	nodegroupv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/containerinfra/nodegroupv1"
	recordsetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/dns/recordsetv2"
	zonev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/dns/zonev2"
	ec2credentialv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/ec2credentialv3"
	projectv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/projectv3"
	roleassignmentv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/roleassignmentv3"
	rolev3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/rolev3"
	userv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/userv3"
	addressscopev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/addressscopev2"
	floatingipv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/floatingipv2"
	networkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/networkv2"
	portv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/portv2"
	routerinterfacev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerinterfacev2"
	routerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerv2"
	subnetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/subnetv2"
	containerv1 "github.com/crossplane-contrib/provider-openstack/internal/controller/objectstorage/containerv1"
	providerconfig "github.com/crossplane-contrib/provider-openstack/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instancev2.Setup,
		keypairv2.Setup,
		clustertemplatev1.Setup,
		clusterv1.Setup,
		nodegroupv1.Setup,
		recordsetv2.Setup,
		zonev2.Setup,
		ec2credentialv3.Setup,
		projectv3.Setup,
		roleassignmentv3.Setup,
		rolev3.Setup,
		userv3.Setup,
		addressscopev2.Setup,
		floatingipv2.Setup,
		networkv2.Setup,
		portv2.Setup,
		routerinterfacev2.Setup,
		routerv2.Setup,
		subnetv2.Setup,
		containerv1.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
