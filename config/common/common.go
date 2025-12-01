/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"fmt"
	"strconv"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/globallogicuki/provider-harbor/config/common"

	// RegistryIDExtractor is the golang path to ExtractAccessor function
	// in this package.
	RegistryIDExtractor = SelfPackagePath + ".ExtractRegistryID()"

	// ProjectNameExtractor is the golang path to ExtractProjectName function
	// in this package.
	ProjectNameExtractor = SelfPackagePath + ".ExtractProjectName()"
)

// ExtractRegistryID extracts the registryId parameter from a registry object
func ExtractRegistryID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetValue("status.atProvider.registryId")
		if err != nil {
			return ""
		}
		// Convert the interface{} to a string based on its actual type, we expect a float64!
		switch v := r.(type) {
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64)
		default:
			fmt.Printf("unsupported type: %T\n", v)
			return ""
		}
	}
}

// ExtractProjectName extracts the name parameter from a project object
func ExtractProjectName() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetValue("status.atProvider.name")
		if err != nil {
			return ""
		}
		// The name field is a string
		if s, ok := r.(string); ok {
			return s
		}
		return ""
	}
}
