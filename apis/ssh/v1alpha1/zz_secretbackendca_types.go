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

type SecretBackendCAObservation struct {

	// The path of the SSH Secret Backend where the CA should be configured
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Whether Vault should generate the signing key pair internally.
	GenerateSigningKey *bool `json:"generateSigningKey,omitempty" tf:"generate_signing_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Public key part the SSH CA key pair; required if generate_signing_key is false.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SecretBackendCAParameters struct {

	// The path of the SSH Secret Backend where the CA should be configured
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Whether Vault should generate the signing key pair internally.
	// +kubebuilder:validation:Optional
	GenerateSigningKey *bool `json:"generateSigningKey,omitempty" tf:"generate_signing_key,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Private key part the SSH CA key pair; required if generate_signing_key is false.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Public key part the SSH CA key pair; required if generate_signing_key is false.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

// SecretBackendCASpec defines the desired state of SecretBackendCA
type SecretBackendCASpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendCAParameters `json:"forProvider"`
}

// SecretBackendCAStatus defines the observed state of SecretBackendCA.
type SecretBackendCAStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendCAObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendCA is the Schema for the SecretBackendCAs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendCA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretBackendCASpec   `json:"spec"`
	Status            SecretBackendCAStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendCAList contains a list of SecretBackendCAs
type SecretBackendCAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendCA `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendCA_Kind             = "SecretBackendCA"
	SecretBackendCA_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendCA_Kind}.String()
	SecretBackendCA_KindAPIVersion   = SecretBackendCA_Kind + "." + CRDGroupVersion.String()
	SecretBackendCA_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendCA_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendCA{}, &SecretBackendCAList{})
}