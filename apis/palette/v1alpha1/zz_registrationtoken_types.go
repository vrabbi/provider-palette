// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegistrationTokenInitParameters struct {

	// (String) A brief description of the registration token.
	// A brief description of the registration token.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// MM-DD format.
	// The expiration date of the registration token in `YYYY-MM-DD` format.
	ExpiryDate *string `json:"expiryDate,omitempty" tf:"expiry_date,omitempty"`

	// (String) The name of the registration token.
	// The name of the registration token.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The unique identifier of the project associated with the registration token.
	// The unique identifier of the project associated with the registration token.
	ProjectUID *string `json:"projectUid,omitempty" tf:"project_uid,omitempty"`

	// (String) The status of the registration token. Allowed values are active or inactive. Default is active.
	// The status of the registration token. Allowed values are `active` or `inactive`. Default is `active`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegistrationTokenObservation struct {

	// (String) A brief description of the registration token.
	// A brief description of the registration token.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// MM-DD format.
	// The expiration date of the registration token in `YYYY-MM-DD` format.
	ExpiryDate *string `json:"expiryDate,omitempty" tf:"expiry_date,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the registration token.
	// The name of the registration token.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The unique identifier of the project associated with the registration token.
	// The unique identifier of the project associated with the registration token.
	ProjectUID *string `json:"projectUid,omitempty" tf:"project_uid,omitempty"`

	// (String) The status of the registration token. Allowed values are active or inactive. Default is active.
	// The status of the registration token. Allowed values are `active` or `inactive`. Default is `active`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String)
	Token *string `json:"token,omitempty" tf:"token,omitempty"`
}

type RegistrationTokenParameters struct {

	// (String) A brief description of the registration token.
	// A brief description of the registration token.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// MM-DD format.
	// The expiration date of the registration token in `YYYY-MM-DD` format.
	// +kubebuilder:validation:Optional
	ExpiryDate *string `json:"expiryDate,omitempty" tf:"expiry_date,omitempty"`

	// (String) The name of the registration token.
	// The name of the registration token.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The unique identifier of the project associated with the registration token.
	// The unique identifier of the project associated with the registration token.
	// +kubebuilder:validation:Optional
	ProjectUID *string `json:"projectUid,omitempty" tf:"project_uid,omitempty"`

	// (String) The status of the registration token. Allowed values are active or inactive. Default is active.
	// The status of the registration token. Allowed values are `active` or `inactive`. Default is `active`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// RegistrationTokenSpec defines the desired state of RegistrationToken
type RegistrationTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistrationTokenParameters `json:"forProvider"`
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
	InitProvider RegistrationTokenInitParameters `json:"initProvider,omitempty"`
}

// RegistrationTokenStatus defines the observed state of RegistrationToken.
type RegistrationTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistrationTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegistrationToken is the Schema for the RegistrationTokens API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type RegistrationToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.expiryDate) || (has(self.initProvider) && has(self.initProvider.expiryDate))",message="spec.forProvider.expiryDate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RegistrationTokenSpec   `json:"spec"`
	Status RegistrationTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistrationTokenList contains a list of RegistrationTokens
type RegistrationTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistrationToken `json:"items"`
}

// Repository type metadata.
var (
	RegistrationToken_Kind             = "RegistrationToken"
	RegistrationToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistrationToken_Kind}.String()
	RegistrationToken_KindAPIVersion   = RegistrationToken_Kind + "." + CRDGroupVersion.String()
	RegistrationToken_GroupVersionKind = CRDGroupVersion.WithKind(RegistrationToken_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistrationToken{}, &RegistrationTokenList{})
}
