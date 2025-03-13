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

type ExportLocationsInitParameters struct {
}

type ExportLocationsObservation struct {
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Preferred *string `json:"preferred,omitempty" tf:"preferred,omitempty"`
}

type ExportLocationsParameters struct {
}

type ShareV2InitParameters struct {

	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// One or more metadata key and value pairs as a dictionary of
	// strings.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the share. Changing this updates the name
	// of the existing share.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share. Changing this
	// creates a new share.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of a share network where the share server exists
	// or will be created. If share_network_id is not set and you provide a snapshot_id,
	// the share_network_id value from the snapshot is used. Changing this creates a new share.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/sharedfilesystem/v1alpha1.SharenetworkV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ShareNetworkID *string `json:"shareNetworkId,omitempty" tf:"share_network_id,omitempty"`

	// Reference to a SharenetworkV2 in sharedfilesystem to populate shareNetworkId.
	// +kubebuilder:validation:Optional
	ShareNetworkIDRef *v1.Reference `json:"shareNetworkIdRef,omitempty" tf:"-"`

	// Selector for a SharenetworkV2 in sharedfilesystem to populate shareNetworkId.
	// +kubebuilder:validation:Optional
	ShareNetworkIDSelector *v1.Selector `json:"shareNetworkIdSelector,omitempty" tf:"-"`

	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto *string `json:"shareProto,omitempty" tf:"share_proto,omitempty"`

	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`

	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`
}

type ShareV2Observation struct {

	// The map of metadata, assigned on the share, which has been
	// explicitly and implicitly added.
	// +mapType=granular
	AllMetadata map[string]*string `json:"allMetadata,omitempty" tf:"all_metadata,omitempty"`

	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of export locations. For example, when a share server
	// has more than one network interface, it can have multiple export locations.
	ExportLocations []ExportLocationsObservation `json:"exportLocations,omitempty" tf:"export_locations,omitempty"`

	// Indicates whether a share has replicas or not.
	HasReplicas *bool `json:"hasReplicas,omitempty" tf:"has_replicas,omitempty"`

	// The share host name.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The unique ID for the Share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// One or more metadata key and value pairs as a dictionary of
	// strings.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the share. Changing this updates the name
	// of the existing share.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The owner of the Share.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share. Changing this
	// creates a new share.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The share replication type.
	ReplicationType *string `json:"replicationType,omitempty" tf:"replication_type,omitempty"`

	// The UUID of a share network where the share server exists
	// or will be created. If share_network_id is not set and you provide a snapshot_id,
	// the share_network_id value from the snapshot is used. Changing this creates a new share.
	ShareNetworkID *string `json:"shareNetworkId,omitempty" tf:"share_network_id,omitempty"`

	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto *string `json:"shareProto,omitempty" tf:"share_proto,omitempty"`

	// The UUID of the share server.
	ShareServerID *string `json:"shareServerId,omitempty" tf:"share_server_id,omitempty"`

	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`

	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`
}

type ShareV2Parameters struct {

	// The share availability zone. Changing this creates a
	// new share.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// One or more metadata key and value pairs as a dictionary of
	// strings.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the share. Changing this updates the name
	// of the existing share.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share. Changing this
	// creates a new share.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of a share network where the share server exists
	// or will be created. If share_network_id is not set and you provide a snapshot_id,
	// the share_network_id value from the snapshot is used. Changing this creates a new share.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/sharedfilesystem/v1alpha1.SharenetworkV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ShareNetworkID *string `json:"shareNetworkId,omitempty" tf:"share_network_id,omitempty"`

	// Reference to a SharenetworkV2 in sharedfilesystem to populate shareNetworkId.
	// +kubebuilder:validation:Optional
	ShareNetworkIDRef *v1.Reference `json:"shareNetworkIdRef,omitempty" tf:"-"`

	// Selector for a SharenetworkV2 in sharedfilesystem to populate shareNetworkId.
	// +kubebuilder:validation:Optional
	ShareNetworkIDSelector *v1.Selector `json:"shareNetworkIdSelector,omitempty" tf:"-"`

	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	// +kubebuilder:validation:Optional
	ShareProto *string `json:"shareProto,omitempty" tf:"share_proto,omitempty"`

	// The share type name. If you omit this parameter, the default
	// share type is used.
	// +kubebuilder:validation:Optional
	ShareType *string `json:"shareType,omitempty" tf:"share_type,omitempty"`

	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`
}

// ShareV2Spec defines the desired state of ShareV2
type ShareV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareV2Parameters `json:"forProvider"`
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
	InitProvider ShareV2InitParameters `json:"initProvider,omitempty"`
}

// ShareV2Status defines the observed state of ShareV2.
type ShareV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ShareV2 is the Schema for the ShareV2s API. Configure a Shared File System share.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ShareV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shareProto) || (has(self.initProvider) && has(self.initProvider.shareProto))",message="spec.forProvider.shareProto is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	Spec   ShareV2Spec   `json:"spec"`
	Status ShareV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareV2List contains a list of ShareV2s
type ShareV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShareV2 `json:"items"`
}

// Repository type metadata.
var (
	ShareV2_Kind             = "ShareV2"
	ShareV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ShareV2_Kind}.String()
	ShareV2_KindAPIVersion   = ShareV2_Kind + "." + CRDGroupVersion.String()
	ShareV2_GroupVersionKind = CRDGroupVersion.WithKind(ShareV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ShareV2{}, &ShareV2List{})
}
