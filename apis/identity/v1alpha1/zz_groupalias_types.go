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

type GroupAliasObservation struct {

	// ID of the group to which this is an alias.
	CanonicalID *string `json:"canonicalId,omitempty" tf:"canonical_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Mount accessor to which this alias belongs to.
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Name of the group alias.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type GroupAliasParameters struct {

	// ID of the group to which this is an alias.
	// +kubebuilder:validation:Optional
	CanonicalID *string `json:"canonicalId,omitempty" tf:"canonical_id,omitempty"`

	// Mount accessor to which this alias belongs to.
	// +kubebuilder:validation:Optional
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Name of the group alias.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// GroupAliasSpec defines the desired state of GroupAlias
type GroupAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupAliasParameters `json:"forProvider"`
}

// GroupAliasStatus defines the observed state of GroupAlias.
type GroupAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupAlias is the Schema for the GroupAliass API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type GroupAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.canonicalId)",message="canonicalId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.mountAccessor)",message="mountAccessor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   GroupAliasSpec   `json:"spec"`
	Status GroupAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupAliasList contains a list of GroupAliass
type GroupAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupAlias `json:"items"`
}

// Repository type metadata.
var (
	GroupAlias_Kind             = "GroupAlias"
	GroupAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupAlias_Kind}.String()
	GroupAlias_KindAPIVersion   = GroupAlias_Kind + "." + CRDGroupVersion.String()
	GroupAlias_GroupVersionKind = CRDGroupVersion.WithKind(GroupAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupAlias{}, &GroupAliasList{})
}