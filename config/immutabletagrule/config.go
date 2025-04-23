package immutabletagrule

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure harbor_immutable_tag_rule resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_immutable_tag_rule", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "ImmutableTagRule"
		r.References["project_id"] = config.Reference{
			TerraformName: "harbor_project",
		}

		// Update any matching or excluding fields to improve the Terraform documentation
		for fieldName, s := range r.TerraformResource.Schema {
			if strings.Contains(fieldName, "_excluding") || strings.Contains(fieldName, "_matching") {
				s.Description = "Use doublestar pattern for path matching."
			}
		}
	})
}
