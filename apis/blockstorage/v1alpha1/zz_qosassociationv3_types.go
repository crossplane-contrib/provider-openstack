/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type QosAssociationV3InitParameters struct {

	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosID *string `json:"qosId,omitempty" tf:"qos_id,omitempty"`

	// The region in which to create the qos association.
	// If omitted, the region argument of the provider is used. Changing
	// this creates a new qos association.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the volume_type to associate.
	// Changing this creates a new qos association.
	VolumeTypeID *string `json:"volumeTypeId,omitempty" tf:"volume_type_id,omitempty"`
}

type QosAssociationV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosID *string `json:"qosId,omitempty" tf:"qos_id,omitempty"`

	// The region in which to create the qos association.
	// If omitted, the region argument of the provider is used. Changing
	// this creates a new qos association.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the volume_type to associate.
	// Changing this creates a new qos association.
	VolumeTypeID *string `json:"volumeTypeId,omitempty" tf:"volume_type_id,omitempty"`
}

type QosAssociationV3Parameters struct {

	// ID of the qos to associate. Changing this creates
	// a new qos association.
	// +kubebuilder:validation:Optional
	QosID *string `json:"qosId,omitempty" tf:"qos_id,omitempty"`

	// The region in which to create the qos association.
	// If omitted, the region argument of the provider is used. Changing
	// this creates a new qos association.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the volume_type to associate.
	// Changing this creates a new qos association.
	// +kubebuilder:validation:Optional
	VolumeTypeID *string `json:"volumeTypeId,omitempty" tf:"volume_type_id,omitempty"`
}

// QosAssociationV3Spec defines the desired state of QosAssociationV3
type QosAssociationV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QosAssociationV3Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider QosAssociationV3InitParameters `json:"initProvider,omitempty"`
}

// QosAssociationV3Status defines the observed state of QosAssociationV3.
type QosAssociationV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QosAssociationV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// QosAssociationV3 is the Schema for the QosAssociationV3s API. Manages a V3 Qos association resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type QosAssociationV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.qosId) || (has(self.initProvider) && has(self.initProvider.qosId))",message="spec.forProvider.qosId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.volumeTypeId) || (has(self.initProvider) && has(self.initProvider.volumeTypeId))",message="spec.forProvider.volumeTypeId is a required parameter"
	Spec   QosAssociationV3Spec   `json:"spec"`
	Status QosAssociationV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QosAssociationV3List contains a list of QosAssociationV3s
type QosAssociationV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QosAssociationV3 `json:"items"`
}

// Repository type metadata.
var (
	QosAssociationV3_Kind             = "QosAssociationV3"
	QosAssociationV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QosAssociationV3_Kind}.String()
	QosAssociationV3_KindAPIVersion   = QosAssociationV3_Kind + "." + CRDGroupVersion.String()
	QosAssociationV3_GroupVersionKind = CRDGroupVersion.WithKind(QosAssociationV3_Kind)
)

func init() {
	SchemeBuilder.Register(&QosAssociationV3{}, &QosAssociationV3List{})
}
