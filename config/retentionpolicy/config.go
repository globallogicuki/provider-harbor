package retentionpolicy

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure harbor_retention_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_retention_policy", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "RetentionPolicy"
		r.References["scope"] = config.Reference{
			TerraformName: "harbor_project",
		}

		// Update any matching or excluding fields to improve the Terraform documentation
		updateMatchingFieldDescriptions(r.TerraformResource.Schema)
	})
}

// updateMatchingFieldDescriptions recursively updates the descriptions of all fields
// with names containing "_excluding" or "_matching" at any nesting level
func updateMatchingFieldDescriptions(schemaMap map[string]*schema.Schema) {
	for fieldName, s := range schemaMap {
		// Check if the current field name matches our criteria
		if strings.Contains(fieldName, "_excluding") || strings.Contains(fieldName, "_matching") {
			s.Description = "Use doublestar pattern for path matching."
		}

		// Check for nested fields in TypeList, TypeSet, or TypeMap with Elem as *schema.Resource
		if elem, ok := s.Elem.(*schema.Resource); ok && elem != nil {
			// Recursive call for nested fields
			updateMatchingFieldDescriptions(elem.Schema)
		}
	}
}
