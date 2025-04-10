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

type PrivateCloudGatewayDNSMapInitParameters struct {

	// (String) The data center in which the private cloud resides.
	// The data center in which the private cloud resides.
	DataCenter *string `json:"dataCenter,omitempty" tf:"data_center,omitempty"`

	// (String) The network to which the private cloud gateway is mapped.
	// The network to which the private cloud gateway is mapped.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) The ID of the Private Cloud Gateway.
	// The ID of the Private Cloud Gateway.
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The domain name used for DNS search queries within the private cloud.
	// The domain name used for DNS search queries within the private cloud.
	SearchDomainName *string `json:"searchDomainName,omitempty" tf:"search_domain_name,omitempty"`
}

type PrivateCloudGatewayDNSMapObservation struct {

	// (String) The data center in which the private cloud resides.
	// The data center in which the private cloud resides.
	DataCenter *string `json:"dataCenter,omitempty" tf:"data_center,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The network to which the private cloud gateway is mapped.
	// The network to which the private cloud gateway is mapped.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) The ID of the Private Cloud Gateway.
	// The ID of the Private Cloud Gateway.
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The domain name used for DNS search queries within the private cloud.
	// The domain name used for DNS search queries within the private cloud.
	SearchDomainName *string `json:"searchDomainName,omitempty" tf:"search_domain_name,omitempty"`
}

type PrivateCloudGatewayDNSMapParameters struct {

	// (String) The data center in which the private cloud resides.
	// The data center in which the private cloud resides.
	// +kubebuilder:validation:Optional
	DataCenter *string `json:"dataCenter,omitempty" tf:"data_center,omitempty"`

	// (String) The network to which the private cloud gateway is mapped.
	// The network to which the private cloud gateway is mapped.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (String) The ID of the Private Cloud Gateway.
	// The ID of the Private Cloud Gateway.
	// +kubebuilder:validation:Optional
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The domain name used for DNS search queries within the private cloud.
	// The domain name used for DNS search queries within the private cloud.
	// +kubebuilder:validation:Optional
	SearchDomainName *string `json:"searchDomainName,omitempty" tf:"search_domain_name,omitempty"`
}

// PrivateCloudGatewayDNSMapSpec defines the desired state of PrivateCloudGatewayDNSMap
type PrivateCloudGatewayDNSMapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateCloudGatewayDNSMapParameters `json:"forProvider"`
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
	InitProvider PrivateCloudGatewayDNSMapInitParameters `json:"initProvider,omitempty"`
}

// PrivateCloudGatewayDNSMapStatus defines the observed state of PrivateCloudGatewayDNSMap.
type PrivateCloudGatewayDNSMapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateCloudGatewayDNSMapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateCloudGatewayDNSMap is the Schema for the PrivateCloudGatewayDNSMaps API. This resource allows for the management of DNS mappings for private cloud gateways. This helps ensure proper DNS resolution for resources within the private cloud environment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type PrivateCloudGatewayDNSMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataCenter) || (has(self.initProvider) && has(self.initProvider.dataCenter))",message="spec.forProvider.dataCenter is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.network) || (has(self.initProvider) && has(self.initProvider.network))",message="spec.forProvider.network is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateCloudGatewayId) || (has(self.initProvider) && has(self.initProvider.privateCloudGatewayId))",message="spec.forProvider.privateCloudGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.searchDomainName) || (has(self.initProvider) && has(self.initProvider.searchDomainName))",message="spec.forProvider.searchDomainName is a required parameter"
	Spec   PrivateCloudGatewayDNSMapSpec   `json:"spec"`
	Status PrivateCloudGatewayDNSMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateCloudGatewayDNSMapList contains a list of PrivateCloudGatewayDNSMaps
type PrivateCloudGatewayDNSMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCloudGatewayDNSMap `json:"items"`
}

// Repository type metadata.
var (
	PrivateCloudGatewayDNSMap_Kind             = "PrivateCloudGatewayDNSMap"
	PrivateCloudGatewayDNSMap_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateCloudGatewayDNSMap_Kind}.String()
	PrivateCloudGatewayDNSMap_KindAPIVersion   = PrivateCloudGatewayDNSMap_Kind + "." + CRDGroupVersion.String()
	PrivateCloudGatewayDNSMap_GroupVersionKind = CRDGroupVersion.WithKind(PrivateCloudGatewayDNSMap_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateCloudGatewayDNSMap{}, &PrivateCloudGatewayDNSMapList{})
}
