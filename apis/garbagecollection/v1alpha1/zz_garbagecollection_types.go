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

type GarbageCollectionInitParameters struct {

	// (Boolean) Allow garbage collection on untagged artifacts.
	DeleteUntagged *bool `json:"deleteUntagged,omitempty" tf:"delete_untagged,omitempty"`

	// (String) Sets the schedule how often the Garbage Collection will run.  Can be to "hourly", "daily", "weekly" or can be a custom cron string ie, "5 4 * * *"
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Number) Number of workers to run the garbage collection, value must be between 1 and 5.
	Workers *float64 `json:"workers,omitempty" tf:"workers,omitempty"`
}

type GarbageCollectionObservation struct {

	// (Boolean) Allow garbage collection on untagged artifacts.
	DeleteUntagged *bool `json:"deleteUntagged,omitempty" tf:"delete_untagged,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Sets the schedule how often the Garbage Collection will run.  Can be to "hourly", "daily", "weekly" or can be a custom cron string ie, "5 4 * * *"
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Number) Number of workers to run the garbage collection, value must be between 1 and 5.
	Workers *float64 `json:"workers,omitempty" tf:"workers,omitempty"`
}

type GarbageCollectionParameters struct {

	// (Boolean) Allow garbage collection on untagged artifacts.
	// +kubebuilder:validation:Optional
	DeleteUntagged *bool `json:"deleteUntagged,omitempty" tf:"delete_untagged,omitempty"`

	// (String) Sets the schedule how often the Garbage Collection will run.  Can be to "hourly", "daily", "weekly" or can be a custom cron string ie, "5 4 * * *"
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// (Number) Number of workers to run the garbage collection, value must be between 1 and 5.
	// +kubebuilder:validation:Optional
	Workers *float64 `json:"workers,omitempty" tf:"workers,omitempty"`
}

// GarbageCollectionSpec defines the desired state of GarbageCollection
type GarbageCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GarbageCollectionParameters `json:"forProvider"`
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
	InitProvider GarbageCollectionInitParameters `json:"initProvider,omitempty"`
}

// GarbageCollectionStatus defines the observed state of GarbageCollection.
type GarbageCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GarbageCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollection is the Schema for the GarbageCollections API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,harbor}
type GarbageCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schedule) || (has(self.initProvider) && has(self.initProvider.schedule))",message="spec.forProvider.schedule is a required parameter"
	Spec   GarbageCollectionSpec   `json:"spec"`
	Status GarbageCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollectionList contains a list of GarbageCollections
type GarbageCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GarbageCollection `json:"items"`
}

// Repository type metadata.
var (
	GarbageCollection_Kind             = "GarbageCollection"
	GarbageCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GarbageCollection_Kind}.String()
	GarbageCollection_KindAPIVersion   = GarbageCollection_Kind + "." + CRDGroupVersion.String()
	GarbageCollection_GroupVersionKind = CRDGroupVersion.WithKind(GarbageCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&GarbageCollection{}, &GarbageCollectionList{})
}