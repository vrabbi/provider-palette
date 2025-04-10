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

type AzureCloudAccountInitParameters struct {

	// (String) Unique client Id from Azure console.
	// Unique client Id from Azure console.
	AzureClientID *string `json:"azureClientId,omitempty" tf:"azure_client_id,omitempty"`

	// (String, Sensitive) Azure secret for authentication.
	// Azure secret for authentication.
	AzureClientSecretSecretRef v1.SecretKeySelector `json:"azureClientSecretSecretRef" tf:"-"`

	// (String) Unique tenant Id from Azure console.
	// Unique tenant Id from Azure console.
	AzureTenantID *string `json:"azureTenantId,omitempty" tf:"azure_tenant_id,omitempty"`

	// (String) The Azure partition in which the cloud account is located.
	// Can be 'AzurePublicCloud' for standard Azure regions or 'AzureUSGovernmentCloud' for Azure GovCloud (US) regions.
	// Default is 'AzurePublicCloud'.
	// The Azure partition in which the cloud account is located.
	// Can be 'AzurePublicCloud' for standard Azure regions or 'AzureUSGovernmentCloud' for Azure GovCloud (US) regions.
	// Default is 'AzurePublicCloud'.
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// (String) The context of the Azure configuration. Defaults to project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Azure configuration. Defaults to `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (Boolean) Disable properties request. This is a boolean value that indicates whether to disable properties request or not. If not specified, the default value is false.
	// Disable properties request. This is a boolean value that indicates whether to disable properties request or not. If not specified, the default value is `false`.
	DisablePropertiesRequest *bool `json:"disablePropertiesRequest,omitempty" tf:"disable_properties_request,omitempty"`

	// (String) The name of the Azure cloud account.
	// The name of the Azure cloud account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The name of the tenant. This is the name of the tenant that is used to connect to the Azure cloud.
	// The name of the tenant. This is the name of the tenant that is used to connect to the Azure cloud.
	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name,omitempty"`
}

type AzureCloudAccountObservation struct {

	// (String) Unique client Id from Azure console.
	// Unique client Id from Azure console.
	AzureClientID *string `json:"azureClientId,omitempty" tf:"azure_client_id,omitempty"`

	// (String) Unique tenant Id from Azure console.
	// Unique tenant Id from Azure console.
	AzureTenantID *string `json:"azureTenantId,omitempty" tf:"azure_tenant_id,omitempty"`

	// (String) The Azure partition in which the cloud account is located.
	// Can be 'AzurePublicCloud' for standard Azure regions or 'AzureUSGovernmentCloud' for Azure GovCloud (US) regions.
	// Default is 'AzurePublicCloud'.
	// The Azure partition in which the cloud account is located.
	// Can be 'AzurePublicCloud' for standard Azure regions or 'AzureUSGovernmentCloud' for Azure GovCloud (US) regions.
	// Default is 'AzurePublicCloud'.
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// (String) The context of the Azure configuration. Defaults to project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Azure configuration. Defaults to `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (Boolean) Disable properties request. This is a boolean value that indicates whether to disable properties request or not. If not specified, the default value is false.
	// Disable properties request. This is a boolean value that indicates whether to disable properties request or not. If not specified, the default value is `false`.
	DisablePropertiesRequest *bool `json:"disablePropertiesRequest,omitempty" tf:"disable_properties_request,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the Azure cloud account.
	// The name of the Azure cloud account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The name of the tenant. This is the name of the tenant that is used to connect to the Azure cloud.
	// The name of the tenant. This is the name of the tenant that is used to connect to the Azure cloud.
	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name,omitempty"`
}

type AzureCloudAccountParameters struct {

	// (String) Unique client Id from Azure console.
	// Unique client Id from Azure console.
	// +kubebuilder:validation:Optional
	AzureClientID *string `json:"azureClientId,omitempty" tf:"azure_client_id,omitempty"`

	// (String, Sensitive) Azure secret for authentication.
	// Azure secret for authentication.
	// +kubebuilder:validation:Optional
	AzureClientSecretSecretRef v1.SecretKeySelector `json:"azureClientSecretSecretRef" tf:"-"`

	// (String) Unique tenant Id from Azure console.
	// Unique tenant Id from Azure console.
	// +kubebuilder:validation:Optional
	AzureTenantID *string `json:"azureTenantId,omitempty" tf:"azure_tenant_id,omitempty"`

	// (String) The Azure partition in which the cloud account is located.
	// Can be 'AzurePublicCloud' for standard Azure regions or 'AzureUSGovernmentCloud' for Azure GovCloud (US) regions.
	// Default is 'AzurePublicCloud'.
	// The Azure partition in which the cloud account is located.
	// Can be 'AzurePublicCloud' for standard Azure regions or 'AzureUSGovernmentCloud' for Azure GovCloud (US) regions.
	// Default is 'AzurePublicCloud'.
	// +kubebuilder:validation:Optional
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// (String) The context of the Azure configuration. Defaults to project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Azure configuration. Defaults to `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (Boolean) Disable properties request. This is a boolean value that indicates whether to disable properties request or not. If not specified, the default value is false.
	// Disable properties request. This is a boolean value that indicates whether to disable properties request or not. If not specified, the default value is `false`.
	// +kubebuilder:validation:Optional
	DisablePropertiesRequest *bool `json:"disablePropertiesRequest,omitempty" tf:"disable_properties_request,omitempty"`

	// (String) The name of the Azure cloud account.
	// The name of the Azure cloud account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// +kubebuilder:validation:Optional
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The name of the tenant. This is the name of the tenant that is used to connect to the Azure cloud.
	// The name of the tenant. This is the name of the tenant that is used to connect to the Azure cloud.
	// +kubebuilder:validation:Optional
	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name,omitempty"`
}

// AzureCloudAccountSpec defines the desired state of AzureCloudAccount
type AzureCloudAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureCloudAccountParameters `json:"forProvider"`
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
	InitProvider AzureCloudAccountInitParameters `json:"initProvider,omitempty"`
}

// AzureCloudAccountStatus defines the observed state of AzureCloudAccount.
type AzureCloudAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureCloudAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AzureCloudAccount is the Schema for the AzureCloudAccounts API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type AzureCloudAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.azureClientId) || (has(self.initProvider) && has(self.initProvider.azureClientId))",message="spec.forProvider.azureClientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.azureClientSecretSecretRef)",message="spec.forProvider.azureClientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.azureTenantId) || (has(self.initProvider) && has(self.initProvider.azureTenantId))",message="spec.forProvider.azureTenantId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AzureCloudAccountSpec   `json:"spec"`
	Status AzureCloudAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureCloudAccountList contains a list of AzureCloudAccounts
type AzureCloudAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureCloudAccount `json:"items"`
}

// Repository type metadata.
var (
	AzureCloudAccount_Kind             = "AzureCloudAccount"
	AzureCloudAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AzureCloudAccount_Kind}.String()
	AzureCloudAccount_KindAPIVersion   = AzureCloudAccount_Kind + "." + CRDGroupVersion.String()
	AzureCloudAccount_GroupVersionKind = CRDGroupVersion.WithKind(AzureCloudAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&AzureCloudAccount{}, &AzureCloudAccountList{})
}
