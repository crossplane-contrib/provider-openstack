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
	projectv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/projectv3"
	roleassignmentv3 "github.com/crossplane-contrib/provider-openstack/internal/controller/identity/roleassignmentv3"
	floatingipv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/floatingipv2"
	networkv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/networkv2"
	routerinterfacev2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerinterfacev2"
	routerv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/routerv2"
	subnetv2 "github.com/crossplane-contrib/provider-openstack/internal/controller/networking/subnetv2"
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
		projectv3.Setup,
		roleassignmentv3.Setup,
		floatingipv2.Setup,
		networkv2.Setup,
		routerinterfacev2.Setup,
		routerv2.Setup,
		subnetv2.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
