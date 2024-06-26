// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/globallogicuki/provider-harbor/apis/group/v1alpha1"
	v1alpha1project "github.com/globallogicuki/provider-harbor/apis/project/v1alpha1"
	v1alpha1projectmembergroup "github.com/globallogicuki/provider-harbor/apis/projectmembergroup/v1alpha1"
	v1alpha1retentionpolicy "github.com/globallogicuki/provider-harbor/apis/retentionpolicy/v1alpha1"
	v1alpha1robotaccount "github.com/globallogicuki/provider-harbor/apis/robotaccount/v1alpha1"
	v1alpha1apis "github.com/globallogicuki/provider-harbor/apis/v1alpha1"
	v1beta1 "github.com/globallogicuki/provider-harbor/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1projectmembergroup.SchemeBuilder.AddToScheme,
		v1alpha1retentionpolicy.SchemeBuilder.AddToScheme,
		v1alpha1robotaccount.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
