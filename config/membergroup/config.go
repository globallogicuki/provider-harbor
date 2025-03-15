package membergroup

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_project_member_group resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_project_member_group", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "MemberGroup"
		r.References["project_id"] = config.Reference{
			TerraformName: "harbor_project",
		}
	})
}
