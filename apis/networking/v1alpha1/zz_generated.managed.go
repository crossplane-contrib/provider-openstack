/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this FloatingipV2.
func (mg *FloatingipV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FloatingipV2.
func (mg *FloatingipV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this FloatingipV2.
func (mg *FloatingipV2) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this FloatingipV2.
func (mg *FloatingipV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FloatingipV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FloatingipV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FloatingipV2.
func (mg *FloatingipV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FloatingipV2.
func (mg *FloatingipV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FloatingipV2.
func (mg *FloatingipV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FloatingipV2.
func (mg *FloatingipV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this FloatingipV2.
func (mg *FloatingipV2) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this FloatingipV2.
func (mg *FloatingipV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FloatingipV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FloatingipV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FloatingipV2.
func (mg *FloatingipV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FloatingipV2.
func (mg *FloatingipV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NetworkV2.
func (mg *NetworkV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkV2.
func (mg *NetworkV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this NetworkV2.
func (mg *NetworkV2) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this NetworkV2.
func (mg *NetworkV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this NetworkV2.
func (mg *NetworkV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this NetworkV2.
func (mg *NetworkV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkV2.
func (mg *NetworkV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkV2.
func (mg *NetworkV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this NetworkV2.
func (mg *NetworkV2) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this NetworkV2.
func (mg *NetworkV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this NetworkV2.
func (mg *NetworkV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this NetworkV2.
func (mg *NetworkV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this QuotaV2.
func (mg *QuotaV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this QuotaV2.
func (mg *QuotaV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this QuotaV2.
func (mg *QuotaV2) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this QuotaV2.
func (mg *QuotaV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this QuotaV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *QuotaV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this QuotaV2.
func (mg *QuotaV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this QuotaV2.
func (mg *QuotaV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this QuotaV2.
func (mg *QuotaV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this QuotaV2.
func (mg *QuotaV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this QuotaV2.
func (mg *QuotaV2) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this QuotaV2.
func (mg *QuotaV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this QuotaV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *QuotaV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this QuotaV2.
func (mg *QuotaV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this QuotaV2.
func (mg *QuotaV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this RouterInterfaceV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *RouterInterfaceV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this RouterInterfaceV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *RouterInterfaceV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this RouterInterfaceV2.
func (mg *RouterInterfaceV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this RouterV2.
func (mg *RouterV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this RouterV2.
func (mg *RouterV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this RouterV2.
func (mg *RouterV2) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this RouterV2.
func (mg *RouterV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this RouterV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *RouterV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this RouterV2.
func (mg *RouterV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this RouterV2.
func (mg *RouterV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this RouterV2.
func (mg *RouterV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this RouterV2.
func (mg *RouterV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this RouterV2.
func (mg *RouterV2) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this RouterV2.
func (mg *RouterV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this RouterV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *RouterV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this RouterV2.
func (mg *RouterV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this RouterV2.
func (mg *RouterV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SubnetV2.
func (mg *SubnetV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SubnetV2.
func (mg *SubnetV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this SubnetV2.
func (mg *SubnetV2) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this SubnetV2.
func (mg *SubnetV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SubnetV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SubnetV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SubnetV2.
func (mg *SubnetV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SubnetV2.
func (mg *SubnetV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SubnetV2.
func (mg *SubnetV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SubnetV2.
func (mg *SubnetV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this SubnetV2.
func (mg *SubnetV2) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this SubnetV2.
func (mg *SubnetV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SubnetV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SubnetV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SubnetV2.
func (mg *SubnetV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SubnetV2.
func (mg *SubnetV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
