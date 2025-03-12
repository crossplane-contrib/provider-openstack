/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ImageAccessV2.
func (mg *ImageAccessV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ImageIDRef,
		Selector:     mg.Spec.ForProvider.ImageIDSelector,
		To: reference.To{
			List:    &ImageV2List{},
			Managed: &ImageV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageID")
	}
	mg.Spec.ForProvider.ImageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ImageID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ImageIDRef,
		Selector:     mg.Spec.InitProvider.ImageIDSelector,
		To: reference.To{
			List:    &ImageV2List{},
			Managed: &ImageV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ImageID")
	}
	mg.Spec.InitProvider.ImageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ImageIDRef = rsp.ResolvedReference

	return nil
}
