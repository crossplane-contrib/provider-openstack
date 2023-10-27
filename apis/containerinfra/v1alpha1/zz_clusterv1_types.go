/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterV1Observation struct {
	APIAddress *string `json:"apiAddress,omitempty" tf:"api_address,omitempty"`

	ClusterTemplateID *string `json:"clusterTemplateId,omitempty" tf:"cluster_template_id,omitempty"`

	CoeVersion *string `json:"coeVersion,omitempty" tf:"coe_version,omitempty"`

	ContainerVersion *string `json:"containerVersion,omitempty" tf:"container_version,omitempty"`

	CreateTimeout *float64 `json:"createTimeout,omitempty" tf:"create_timeout,omitempty"`

	// The time at which node group was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DiscoveryURL *string `json:"discoveryUrl,omitempty" tf:"discovery_url,omitempty"`

	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	FixedNetwork *string `json:"fixedNetwork,omitempty" tf:"fixed_network,omitempty"`

	FixedSubnet *string `json:"fixedSubnet,omitempty" tf:"fixed_subnet,omitempty"`

	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	MasterAddresses []*string `json:"masterAddresses,omitempty" tf:"master_addresses,omitempty"`

	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	MergeLabels *bool `json:"mergeLabels,omitempty" tf:"merge_labels,omitempty"`

	// The name of the node group. Changing this creates a new
	// node group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NodeAddresses []*string `json:"nodeAddresses,omitempty" tf:"node_addresses,omitempty"`

	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The project of the node group. Required if admin
	// wants to create a cluster in another project. Changing this creates a new
	// node group.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the region argument of the provider is used. Changing this creates a new
	// node group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// The time at which node group was created.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type ClusterV1Parameters struct {

	// +kubebuilder:validation:Optional
	ClusterTemplateID *string `json:"clusterTemplateId,omitempty" tf:"cluster_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	CreateTimeout *float64 `json:"createTimeout,omitempty" tf:"create_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	DiscoveryURL *string `json:"discoveryUrl,omitempty" tf:"discovery_url,omitempty"`

	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	// +kubebuilder:validation:Optional
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.NetworkV2
	// +kubebuilder:validation:Optional
	FixedNetwork *string `json:"fixedNetwork,omitempty" tf:"fixed_network,omitempty"`

	// Reference to a NetworkV2 in networking to populate fixedNetwork.
	// +kubebuilder:validation:Optional
	FixedNetworkRef *v1.Reference `json:"fixedNetworkRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate fixedNetwork.
	// +kubebuilder:validation:Optional
	FixedNetworkSelector *v1.Selector `json:"fixedNetworkSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	// +kubebuilder:validation:Optional
	FixedSubnet *string `json:"fixedSubnet,omitempty" tf:"fixed_subnet,omitempty"`

	// Reference to a SubnetV2 in networking to populate fixedSubnet.
	// +kubebuilder:validation:Optional
	FixedSubnetRef *v1.Reference `json:"fixedSubnetRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate fixedSubnet.
	// +kubebuilder:validation:Optional
	FixedSubnetSelector *v1.Selector `json:"fixedSubnetSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// +kubebuilder:validation:Optional
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	// +kubebuilder:validation:Optional
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	// +kubebuilder:validation:Optional
	MergeLabels *bool `json:"mergeLabels,omitempty" tf:"merge_labels,omitempty"`

	// The name of the node group. Changing this creates a new
	// node group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the region argument of the provider is used. Changing this creates a new
	// node group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// ClusterV1Spec defines the desired state of ClusterV1
type ClusterV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterV1Parameters `json:"forProvider"`
}

// ClusterV1Status defines the observed state of ClusterV1.
type ClusterV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterV1 is the Schema for the ClusterV1s API. Manages a V1 Magnum node group resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ClusterV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.clusterTemplateId)",message="clusterTemplateId is a required parameter"
	Spec   ClusterV1Spec   `json:"spec"`
	Status ClusterV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterV1List contains a list of ClusterV1s
type ClusterV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterV1 `json:"items"`
}

// Repository type metadata.
var (
	ClusterV1_Kind             = "ClusterV1"
	ClusterV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterV1_Kind}.String()
	ClusterV1_KindAPIVersion   = ClusterV1_Kind + "." + CRDGroupVersion.String()
	ClusterV1_GroupVersionKind = CRDGroupVersion.WithKind(ClusterV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterV1{}, &ClusterV1List{})
}
