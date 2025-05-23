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

type NodegroupV1InitParameters struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	MergeLabels *bool `json:"mergeLabels,omitempty" tf:"merge_labels,omitempty"`

	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type NodegroupV1Observation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	MergeLabels *bool `json:"mergeLabels,omitempty" tf:"merge_labels,omitempty"`

	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type NodegroupV1Parameters struct {

	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// +kubebuilder:validation:Optional
	MergeLabels *bool `json:"mergeLabels,omitempty" tf:"merge_labels,omitempty"`

	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// NodegroupV1Spec defines the desired state of NodegroupV1
type NodegroupV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodegroupV1Parameters `json:"forProvider"`
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
	InitProvider NodegroupV1InitParameters `json:"initProvider,omitempty"`
}

// NodegroupV1Status defines the observed state of NodegroupV1.
type NodegroupV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodegroupV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodegroupV1 is the Schema for the NodegroupV1s API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type NodegroupV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterId) || (has(self.initProvider) && has(self.initProvider.clusterId))",message="spec.forProvider.clusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NodegroupV1Spec   `json:"spec"`
	Status NodegroupV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodegroupV1List contains a list of NodegroupV1s
type NodegroupV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodegroupV1 `json:"items"`
}

// Repository type metadata.
var (
	NodegroupV1_Kind             = "NodegroupV1"
	NodegroupV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodegroupV1_Kind}.String()
	NodegroupV1_KindAPIVersion   = NodegroupV1_Kind + "." + CRDGroupVersion.String()
	NodegroupV1_GroupVersionKind = CRDGroupVersion.WithKind(NodegroupV1_Kind)
)

func init() {
	SchemeBuilder.Register(&NodegroupV1{}, &NodegroupV1List{})
}
