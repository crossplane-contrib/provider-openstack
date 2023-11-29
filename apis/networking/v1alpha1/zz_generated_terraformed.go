// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this FloatingipV2
func (mg *FloatingipV2) GetTerraformResourceType() string {
	return "openstack_networking_floatingip_v2"
}

// GetConnectionDetailsMapping for this FloatingipV2
func (tr *FloatingipV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FloatingipV2
func (tr *FloatingipV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FloatingipV2
func (tr *FloatingipV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FloatingipV2
func (tr *FloatingipV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FloatingipV2
func (tr *FloatingipV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FloatingipV2
func (tr *FloatingipV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this FloatingipV2
func (tr *FloatingipV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this FloatingipV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FloatingipV2) LateInitialize(attrs []byte) (bool, error) {
	params := &FloatingipV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FloatingipV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NetworkV2
func (mg *NetworkV2) GetTerraformResourceType() string {
	return "openstack_networking_network_v2"
}

// GetConnectionDetailsMapping for this NetworkV2
func (tr *NetworkV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NetworkV2
func (tr *NetworkV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NetworkV2
func (tr *NetworkV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NetworkV2
func (tr *NetworkV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NetworkV2
func (tr *NetworkV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NetworkV2
func (tr *NetworkV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this NetworkV2
func (tr *NetworkV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this NetworkV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NetworkV2) LateInitialize(attrs []byte) (bool, error) {
	params := &NetworkV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NetworkV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this QuotaV2
func (mg *QuotaV2) GetTerraformResourceType() string {
	return "openstack_networking_quota_v2"
}

// GetConnectionDetailsMapping for this QuotaV2
func (tr *QuotaV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this QuotaV2
func (tr *QuotaV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this QuotaV2
func (tr *QuotaV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this QuotaV2
func (tr *QuotaV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this QuotaV2
func (tr *QuotaV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this QuotaV2
func (tr *QuotaV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this QuotaV2
func (tr *QuotaV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this QuotaV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *QuotaV2) LateInitialize(attrs []byte) (bool, error) {
	params := &QuotaV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *QuotaV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RbacPolicyV2
func (mg *RbacPolicyV2) GetTerraformResourceType() string {
	return "openstack_networking_rbac_policy_v2"
}

// GetConnectionDetailsMapping for this RbacPolicyV2
func (tr *RbacPolicyV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RbacPolicyV2
func (tr *RbacPolicyV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RbacPolicyV2
func (tr *RbacPolicyV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RbacPolicyV2
func (tr *RbacPolicyV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RbacPolicyV2
func (tr *RbacPolicyV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RbacPolicyV2
func (tr *RbacPolicyV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this RbacPolicyV2
func (tr *RbacPolicyV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this RbacPolicyV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RbacPolicyV2) LateInitialize(attrs []byte) (bool, error) {
	params := &RbacPolicyV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RbacPolicyV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RouterInterfaceV2
func (mg *RouterInterfaceV2) GetTerraformResourceType() string {
	return "openstack_networking_router_interface_v2"
}

// GetConnectionDetailsMapping for this RouterInterfaceV2
func (tr *RouterInterfaceV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RouterInterfaceV2
func (tr *RouterInterfaceV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RouterInterfaceV2
func (tr *RouterInterfaceV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RouterInterfaceV2
func (tr *RouterInterfaceV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RouterInterfaceV2
func (tr *RouterInterfaceV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RouterInterfaceV2
func (tr *RouterInterfaceV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this RouterInterfaceV2
func (tr *RouterInterfaceV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this RouterInterfaceV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RouterInterfaceV2) LateInitialize(attrs []byte) (bool, error) {
	params := &RouterInterfaceV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RouterInterfaceV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RouterV2
func (mg *RouterV2) GetTerraformResourceType() string {
	return "openstack_networking_router_v2"
}

// GetConnectionDetailsMapping for this RouterV2
func (tr *RouterV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RouterV2
func (tr *RouterV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RouterV2
func (tr *RouterV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RouterV2
func (tr *RouterV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RouterV2
func (tr *RouterV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RouterV2
func (tr *RouterV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this RouterV2
func (tr *RouterV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this RouterV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RouterV2) LateInitialize(attrs []byte) (bool, error) {
	params := &RouterV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("ExternalGateway"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RouterV2) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubnetV2
func (mg *SubnetV2) GetTerraformResourceType() string {
	return "openstack_networking_subnet_v2"
}

// GetConnectionDetailsMapping for this SubnetV2
func (tr *SubnetV2) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubnetV2
func (tr *SubnetV2) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubnetV2
func (tr *SubnetV2) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubnetV2
func (tr *SubnetV2) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubnetV2
func (tr *SubnetV2) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubnetV2
func (tr *SubnetV2) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SubnetV2
func (tr *SubnetV2) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SubnetV2 using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubnetV2) LateInitialize(attrs []byte) (bool, error) {
	params := &SubnetV2Parameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("AllocationPools"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubnetV2) GetTerraformSchemaVersion() int {
	return 0
}
