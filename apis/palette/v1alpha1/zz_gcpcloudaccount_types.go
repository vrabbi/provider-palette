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

type GCPCloudAccountInitParameters struct {

	// (String) The context of the GCP configuration. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the GCP configuration. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String, Sensitive) The GCP credentials in JSON format. These credentials are required to authenticate and manage.
	// The GCP credentials in JSON format. These credentials are required to authenticate and manage.
	GCPJSONCredentials *string `json:"gcpJsonCredentials,omitempty" tf:"gcp_json_credentials,omitempty"`

	// (String) The name of the GCP account.
	// The name of the GCP account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GCPCloudAccountObservation struct {

	// (String) The context of the GCP configuration. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the GCP configuration. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String, Sensitive) The GCP credentials in JSON format. These credentials are required to authenticate and manage.
	// The GCP credentials in JSON format. These credentials are required to authenticate and manage.
	GCPJSONCredentials *string `json:"gcpJsonCredentials,omitempty" tf:"gcp_json_credentials,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the GCP account.
	// The name of the GCP account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GCPCloudAccountParameters struct {

	// (String) The context of the GCP configuration. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the GCP configuration. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String, Sensitive) The GCP credentials in JSON format. These credentials are required to authenticate and manage.
	// The GCP credentials in JSON format. These credentials are required to authenticate and manage.
	// +kubebuilder:validation:Optional
	GCPJSONCredentials *string `json:"gcpJsonCredentials,omitempty" tf:"gcp_json_credentials,omitempty"`

	// (String) The name of the GCP account.
	// The name of the GCP account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// GCPCloudAccountSpec defines the desired state of GCPCloudAccount
type GCPCloudAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GCPCloudAccountParameters `json:"forProvider"`
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
	InitProvider GCPCloudAccountInitParameters `json:"initProvider,omitempty"`
}

// GCPCloudAccountStatus defines the observed state of GCPCloudAccount.
type GCPCloudAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GCPCloudAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GCPCloudAccount is the Schema for the GCPCloudAccounts API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type GCPCloudAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gcpJsonCredentials) || (has(self.initProvider) && has(self.initProvider.gcpJsonCredentials))",message="spec.forProvider.gcpJsonCredentials is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GCPCloudAccountSpec   `json:"spec"`
	Status GCPCloudAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPCloudAccountList contains a list of GCPCloudAccounts
type GCPCloudAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPCloudAccount `json:"items"`
}

// Repository type metadata.
var (
	GCPCloudAccount_Kind             = "GCPCloudAccount"
	GCPCloudAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GCPCloudAccount_Kind}.String()
	GCPCloudAccount_KindAPIVersion   = GCPCloudAccount_Kind + "." + CRDGroupVersion.String()
	GCPCloudAccount_GroupVersionKind = CRDGroupVersion.WithKind(GCPCloudAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&GCPCloudAccount{}, &GCPCloudAccountList{})
}
