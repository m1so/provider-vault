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

type AuthBackendObservation struct {

	// The accessor of the JWT auth backend
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	// The value against which to match the iss claim in a JWT
	BoundIssuer *string `json:"boundIssuer,omitempty" tf:"bound_issuer,omitempty"`

	// The default role to use if none is provided during login
	DefaultRole *string `json:"defaultRole,omitempty" tf:"default_role,omitempty"`

	// The description of the auth backend
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCAPem *string `json:"jwksCaPem,omitempty" tf:"jwks_ca_pem,omitempty"`

	// JWKS URL to use to authenticate signatures. Cannot be used with 'oidc_discovery_url' or 'jwt_validation_pubkeys'.
	JwksURL *string `json:"jwksUrl,omitempty" tf:"jwks_url,omitempty"`

	// A list of supported signing algorithms. Defaults to [RS256]
	JwtSupportedAlgs []*string `json:"jwtSupportedAlgs,omitempty" tf:"jwt_supported_algs,omitempty"`

	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with 'jwks_url' or 'oidc_discovery_url'.
	JwtValidationPubkeys []*string `json:"jwtValidationPubkeys,omitempty" tf:"jwt_validation_pubkeys,omitempty"`

	// Specifies if the auth method is local only
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs.
	NamespaceInState *bool `json:"namespaceInState,omitempty" tf:"namespace_in_state,omitempty"`

	// Client ID used for OIDC
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCAPem *string `json:"oidcDiscoveryCaPem,omitempty" tf:"oidc_discovery_ca_pem,omitempty"`

	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used with 'jwks_url' or 'jwt_validation_pubkeys'.
	OidcDiscoveryURL *string `json:"oidcDiscoveryUrl,omitempty" tf:"oidc_discovery_url,omitempty"`

	// The response mode to be used in the OAuth2 request. Allowed values are 'query' and 'form_post'. Defaults to 'query'. If using Vault namespaces, and oidc_response_mode is 'form_post', then 'namespace_in_state' should be set to false.
	OidcResponseMode *string `json:"oidcResponseMode,omitempty" tf:"oidc_response_mode,omitempty"`

	// The response types to request. Allowed values are 'code' and 'id_token'. Defaults to 'code'. Note: 'id_token' may only be used if 'oidc_response_mode' is set to 'form_post'.
	OidcResponseTypes []*string `json:"oidcResponseTypes,omitempty" tf:"oidc_response_types,omitempty"`

	// path to mount the backend
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Provider specific handling configuration
	ProviderConfig map[string]*string `json:"providerConfig,omitempty" tf:"provider_config,omitempty"`

	Tune []TuneObservation `json:"tune,omitempty" tf:"tune,omitempty"`

	// Type of backend. Can be either 'jwt' or 'oidc'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthBackendParameters struct {

	// The value against which to match the iss claim in a JWT
	// +kubebuilder:validation:Optional
	BoundIssuer *string `json:"boundIssuer,omitempty" tf:"bound_issuer,omitempty"`

	// The default role to use if none is provided during login
	// +kubebuilder:validation:Optional
	DefaultRole *string `json:"defaultRole,omitempty" tf:"default_role,omitempty"`

	// The description of the auth backend
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	// +kubebuilder:validation:Optional
	JwksCAPem *string `json:"jwksCaPem,omitempty" tf:"jwks_ca_pem,omitempty"`

	// JWKS URL to use to authenticate signatures. Cannot be used with 'oidc_discovery_url' or 'jwt_validation_pubkeys'.
	// +kubebuilder:validation:Optional
	JwksURL *string `json:"jwksUrl,omitempty" tf:"jwks_url,omitempty"`

	// A list of supported signing algorithms. Defaults to [RS256]
	// +kubebuilder:validation:Optional
	JwtSupportedAlgs []*string `json:"jwtSupportedAlgs,omitempty" tf:"jwt_supported_algs,omitempty"`

	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with 'jwks_url' or 'oidc_discovery_url'.
	// +kubebuilder:validation:Optional
	JwtValidationPubkeys []*string `json:"jwtValidationPubkeys,omitempty" tf:"jwt_validation_pubkeys,omitempty"`

	// Specifies if the auth method is local only
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs.
	// +kubebuilder:validation:Optional
	NamespaceInState *bool `json:"namespaceInState,omitempty" tf:"namespace_in_state,omitempty"`

	// Client ID used for OIDC
	// +kubebuilder:validation:Optional
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// Client Secret used for OIDC
	// +kubebuilder:validation:Optional
	OidcClientSecretSecretRef *v1.SecretKeySelector `json:"oidcClientSecretSecretRef,omitempty" tf:"-"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	// +kubebuilder:validation:Optional
	OidcDiscoveryCAPem *string `json:"oidcDiscoveryCaPem,omitempty" tf:"oidc_discovery_ca_pem,omitempty"`

	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used with 'jwks_url' or 'jwt_validation_pubkeys'.
	// +kubebuilder:validation:Optional
	OidcDiscoveryURL *string `json:"oidcDiscoveryUrl,omitempty" tf:"oidc_discovery_url,omitempty"`

	// The response mode to be used in the OAuth2 request. Allowed values are 'query' and 'form_post'. Defaults to 'query'. If using Vault namespaces, and oidc_response_mode is 'form_post', then 'namespace_in_state' should be set to false.
	// +kubebuilder:validation:Optional
	OidcResponseMode *string `json:"oidcResponseMode,omitempty" tf:"oidc_response_mode,omitempty"`

	// The response types to request. Allowed values are 'code' and 'id_token'. Defaults to 'code'. Note: 'id_token' may only be used if 'oidc_response_mode' is set to 'form_post'.
	// +kubebuilder:validation:Optional
	OidcResponseTypes []*string `json:"oidcResponseTypes,omitempty" tf:"oidc_response_types,omitempty"`

	// path to mount the backend
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Provider specific handling configuration
	// +kubebuilder:validation:Optional
	ProviderConfig map[string]*string `json:"providerConfig,omitempty" tf:"provider_config,omitempty"`

	// +kubebuilder:validation:Optional
	Tune []TuneParameters `json:"tune,omitempty" tf:"tune,omitempty"`

	// Type of backend. Can be either 'jwt' or 'oidc'
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TuneObservation struct {
	AllowedResponseHeaders []*string `json:"allowedResponseHeaders,omitempty" tf:"allowed_response_headers,omitempty"`

	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys,omitempty"`

	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys,omitempty"`

	DefaultLeaseTTL *string `json:"defaultLeaseTtl,omitempty" tf:"default_lease_ttl,omitempty"`

	ListingVisibility *string `json:"listingVisibility,omitempty" tf:"listing_visibility,omitempty"`

	MaxLeaseTTL *string `json:"maxLeaseTtl,omitempty" tf:"max_lease_ttl,omitempty"`

	PassthroughRequestHeaders []*string `json:"passthroughRequestHeaders,omitempty" tf:"passthrough_request_headers,omitempty"`

	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type TuneParameters struct {

	// +kubebuilder:validation:Optional
	AllowedResponseHeaders []*string `json:"allowedResponseHeaders,omitempty" tf:"allowed_response_headers"`

	// +kubebuilder:validation:Optional
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys"`

	// +kubebuilder:validation:Optional
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys"`

	// +kubebuilder:validation:Optional
	DefaultLeaseTTL *string `json:"defaultLeaseTtl,omitempty" tf:"default_lease_ttl"`

	// +kubebuilder:validation:Optional
	ListingVisibility *string `json:"listingVisibility,omitempty" tf:"listing_visibility"`

	// +kubebuilder:validation:Optional
	MaxLeaseTTL *string `json:"maxLeaseTtl,omitempty" tf:"max_lease_ttl"`

	// +kubebuilder:validation:Optional
	PassthroughRequestHeaders []*string `json:"passthroughRequestHeaders,omitempty" tf:"passthrough_request_headers"`

	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type"`
}

// AuthBackendSpec defines the desired state of AuthBackend
type AuthBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendParameters `json:"forProvider"`
}

// AuthBackendStatus defines the observed state of AuthBackend.
type AuthBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackend is the Schema for the AuthBackends API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendSpec   `json:"spec"`
	Status            AuthBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendList contains a list of AuthBackends
type AuthBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackend `json:"items"`
}

// Repository type metadata.
var (
	AuthBackend_Kind             = "AuthBackend"
	AuthBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackend_Kind}.String()
	AuthBackend_KindAPIVersion   = AuthBackend_Kind + "." + CRDGroupVersion.String()
	AuthBackend_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackend{}, &AuthBackendList{})
}