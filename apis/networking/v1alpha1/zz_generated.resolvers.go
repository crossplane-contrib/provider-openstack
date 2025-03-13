/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &PortV2List{},
			Managed: &PortV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PortID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.PortIDRef,
		Selector:     mg.Spec.InitProvider.PortIDSelector,
		To: reference.To{
			List:    &PortV2List{},
			Managed: &PortV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PortID")
	}
	mg.Spec.InitProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PortIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PortV2.
func (mg *PortV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.FixedIP); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FixedIP[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.FixedIP[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.FixedIP[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetV2List{},
				Managed: &SubnetV2{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.FixedIP[i3].SubnetID")
		}
		mg.Spec.ForProvider.FixedIP[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.FixedIP[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkV2List{},
			Managed: &NetworkV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.FixedIP); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FixedIP[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.FixedIP[i3].SubnetIDRef,
			Selector:     mg.Spec.InitProvider.FixedIP[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetV2List{},
				Managed: &SubnetV2{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.FixedIP[i3].SubnetID")
		}
		mg.Spec.InitProvider.FixedIP[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.FixedIP[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkV2List{},
			Managed: &NetworkV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QosBandwidthLimitRuleV2.
func (mg *QosBandwidthLimitRuleV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QosPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.QosPolicyIDRef,
		Selector:     mg.Spec.ForProvider.QosPolicyIDSelector,
		To: reference.To{
			List:    &QosPolicyV2List{},
			Managed: &QosPolicyV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QosPolicyID")
	}
	mg.Spec.ForProvider.QosPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QosPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QosPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.QosPolicyIDRef,
		Selector:     mg.Spec.InitProvider.QosPolicyIDSelector,
		To: reference.To{
			List:    &QosPolicyV2List{},
			Managed: &QosPolicyV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QosPolicyID")
	}
	mg.Spec.InitProvider.QosPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QosPolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QosDscpMarkingRuleV2.
func (mg *QosDscpMarkingRuleV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QosPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.QosPolicyIDRef,
		Selector:     mg.Spec.ForProvider.QosPolicyIDSelector,
		To: reference.To{
			List:    &QosPolicyV2List{},
			Managed: &QosPolicyV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QosPolicyID")
	}
	mg.Spec.ForProvider.QosPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QosPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QosPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.QosPolicyIDRef,
		Selector:     mg.Spec.InitProvider.QosPolicyIDSelector,
		To: reference.To{
			List:    &QosPolicyV2List{},
			Managed: &QosPolicyV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QosPolicyID")
	}
	mg.Spec.InitProvider.QosPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QosPolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QosMinimumBandwidthRuleV2.
func (mg *QosMinimumBandwidthRuleV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QosPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.QosPolicyIDRef,
		Selector:     mg.Spec.ForProvider.QosPolicyIDSelector,
		To: reference.To{
			List:    &QosPolicyV2List{},
			Managed: &QosPolicyV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QosPolicyID")
	}
	mg.Spec.ForProvider.QosPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QosPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QosPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.QosPolicyIDRef,
		Selector:     mg.Spec.InitProvider.QosPolicyIDSelector,
		To: reference.To{
			List:    &QosPolicyV2List{},
			Managed: &QosPolicyV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QosPolicyID")
	}
	mg.Spec.InitProvider.QosPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QosPolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QuotaV2.
func (mg *QuotaV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectIDRef,
		Selector:     mg.Spec.ForProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectV3List{},
			Managed: &v1alpha1.ProjectV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectID")
	}
	mg.Spec.ForProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectIDRef,
		Selector:     mg.Spec.InitProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectV3List{},
			Managed: &v1alpha1.ProjectV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProjectID")
	}
	mg.Spec.InitProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RbacPolicyV2.
func (mg *RbacPolicyV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ObjectID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ObjectIDRef,
		Selector:     mg.Spec.ForProvider.ObjectIDSelector,
		To: reference.To{
			List:    &NetworkV2List{},
			Managed: &NetworkV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ObjectID")
	}
	mg.Spec.ForProvider.ObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ObjectID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ObjectIDRef,
		Selector:     mg.Spec.InitProvider.ObjectIDSelector,
		To: reference.To{
			List:    &NetworkV2List{},
			Managed: &NetworkV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ObjectID")
	}
	mg.Spec.InitProvider.ObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ObjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouterIDRef,
		Selector:     mg.Spec.ForProvider.RouterIDSelector,
		To: reference.To{
			List:    &RouterV2List{},
			Managed: &RouterV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouterID")
	}
	mg.Spec.ForProvider.RouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetV2List{},
			Managed: &SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RouterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RouterIDRef,
		Selector:     mg.Spec.InitProvider.RouterIDSelector,
		To: reference.To{
			List:    &RouterV2List{},
			Managed: &RouterV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RouterID")
	}
	mg.Spec.InitProvider.RouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RouterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetV2List{},
			Managed: &SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouterRouteV2.
func (mg *RouterRouteV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouterID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RouterIDRef,
		Selector:     mg.Spec.ForProvider.RouterIDSelector,
		To: reference.To{
			List:    &RouterV2List{},
			Managed: &RouterV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouterID")
	}
	mg.Spec.ForProvider.RouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RouterID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.RouterIDRef,
		Selector:     mg.Spec.InitProvider.RouterIDSelector,
		To: reference.To{
			List:    &RouterV2List{},
			Managed: &RouterV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RouterID")
	}
	mg.Spec.InitProvider.RouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RouterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecgroupRuleV2.
func (mg *SecgroupRuleV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecgroupV2List{},
			Managed: &SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecgroupV2List{},
			Managed: &SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetRouteV2.
func (mg *SubnetRouteV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetV2List{},
			Managed: &SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetV2List{},
			Managed: &SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetV2.
func (mg *SubnetV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &SubnetV2List{},
			Managed: &SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &SubnetV2List{},
			Managed: &SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TrunkV2.
func (mg *TrunkV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &PortV2List{},
			Managed: &PortV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SubPort); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubPort[i3].PortID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubPort[i3].PortIDRef,
			Selector:     mg.Spec.ForProvider.SubPort[i3].PortIDSelector,
			To: reference.To{
				List:    &PortV2List{},
				Managed: &PortV2{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SubPort[i3].PortID")
		}
		mg.Spec.ForProvider.SubPort[i3].PortID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SubPort[i3].PortIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PortID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.PortIDRef,
		Selector:     mg.Spec.InitProvider.PortIDSelector,
		To: reference.To{
			List:    &PortV2List{},
			Managed: &PortV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PortID")
	}
	mg.Spec.InitProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PortIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.SubPort); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubPort[i3].PortID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubPort[i3].PortIDRef,
			Selector:     mg.Spec.InitProvider.SubPort[i3].PortIDSelector,
			To: reference.To{
				List:    &PortV2List{},
				Managed: &PortV2{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SubPort[i3].PortID")
		}
		mg.Spec.InitProvider.SubPort[i3].PortID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SubPort[i3].PortIDRef = rsp.ResolvedReference

	}

	return nil
}
