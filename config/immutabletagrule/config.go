package immutabletagrule

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_immutable_tag_rule resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_immutable_tag_rule", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "ImmutableTagRule"
		r.References["project_id"] = config.Reference{
			Type: "github.com/globallogicuki/provider-harbor/apis/project/v1alpha1.Project",
		}
	})
}
