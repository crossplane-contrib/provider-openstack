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

// ResolveReferences of this QosAssociationV3.
func (mg *QosAssociationV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QosID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.QosIDRef,
		Selector:     mg.Spec.ForProvider.QosIDSelector,
		To: reference.To{
			List:    &QosV3List{},
			Managed: &QosV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QosID")
	}
	mg.Spec.ForProvider.QosID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QosIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeTypeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VolumeTypeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeTypeIDSelector,
		To: reference.To{
			List:    &VolumeTypeV3List{},
			Managed: &VolumeTypeV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeTypeID")
	}
	mg.Spec.ForProvider.VolumeTypeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeTypeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QosID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.QosIDRef,
		Selector:     mg.Spec.InitProvider.QosIDSelector,
		To: reference.To{
			List:    &QosV3List{},
			Managed: &QosV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QosID")
	}
	mg.Spec.InitProvider.QosID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QosIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VolumeTypeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.VolumeTypeIDRef,
		Selector:     mg.Spec.InitProvider.VolumeTypeIDSelector,
		To: reference.To{
			List:    &VolumeTypeV3List{},
			Managed: &VolumeTypeV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VolumeTypeID")
	}
	mg.Spec.InitProvider.VolumeTypeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VolumeTypeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QuotasetV3.
func (mg *QuotasetV3) ResolveReferences(ctx context.Context, c client.Reader) error {
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

// ResolveReferences of this VolumeAttachV3.
func (mg *VolumeAttachV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeV3List{},
			Managed: &VolumeV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VolumeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.VolumeIDRef,
		Selector:     mg.Spec.InitProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeV3List{},
			Managed: &VolumeV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VolumeID")
	}
	mg.Spec.InitProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VolumeTypeAccessV3.
func (mg *VolumeTypeAccessV3) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectID),
		Extract:      resource.ExtractResourceID(),
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
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeTypeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VolumeTypeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeTypeIDSelector,
		To: reference.To{
			List:    &VolumeTypeV3List{},
			Managed: &VolumeTypeV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeTypeID")
	}
	mg.Spec.ForProvider.VolumeTypeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeTypeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectID),
		Extract:      resource.ExtractResourceID(),
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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VolumeTypeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.VolumeTypeIDRef,
		Selector:     mg.Spec.InitProvider.VolumeTypeIDSelector,
		To: reference.To{
			List:    &VolumeTypeV3List{},
			Managed: &VolumeTypeV3{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VolumeTypeID")
	}
	mg.Spec.InitProvider.VolumeTypeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VolumeTypeIDRef = rsp.ResolvedReference

	return nil
}
