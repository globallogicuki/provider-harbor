package memberuser

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_project_member_user resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_project_member_user", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "MemberUser"
		r.References["project_id"] = config.Reference{
			TerraformName: "harbor_project",
		}
	})
}
