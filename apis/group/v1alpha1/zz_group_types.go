// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type GroupInitParameters struct {

	// (String) The name of the group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// (Number) 3. Note: group type 3 is OIDC group.
	GroupType *float64 `json:"groupType,omitempty" tf:"group_type,omitempty"`

	// (String) The distinguished name of the group within AD/LDAP.
	LdapGroupDn *string `json:"ldapGroupDn,omitempty" tf:"ldap_group_dn,omitempty"`
}

type GroupObservation struct {

	// (String) The name of the group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// (Number) 3. Note: group type 3 is OIDC group.
	GroupType *float64 `json:"groupType,omitempty" tf:"group_type,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The distinguished name of the group within AD/LDAP.
	LdapGroupDn *string `json:"ldapGroupDn,omitempty" tf:"ldap_group_dn,omitempty"`
}

type GroupParameters struct {

	// (String) The name of the group.
	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// (Number) 3. Note: group type 3 is OIDC group.
	// +kubebuilder:validation:Optional
	GroupType *float64 `json:"groupType,omitempty" tf:"group_type,omitempty"`

	// (String) The distinguished name of the group within AD/LDAP.
	// +kubebuilder:validation:Optional
	LdapGroupDn *string `json:"ldapGroupDn,omitempty" tf:"ldap_group_dn,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,harbor}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupName) || (has(self.initProvider) && has(self.initProvider.groupName))",message="spec.forProvider.groupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupType) || (has(self.initProvider) && has(self.initProvider.groupType))",message="spec.forProvider.groupType is a required parameter"
	Spec   GroupSpec   `json:"spec"`
	Status GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
