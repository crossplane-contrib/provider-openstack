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
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this InstanceV2.
func (mg *InstanceV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &KeypairV2List{},
			Managed: &KeypairV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupsSelector,
		To: reference.To{
			List:    &SecgroupV2List{},
			Managed: &SecgroupV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this QuotasetV2.
func (mg *QuotasetV2) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	return nil
}
