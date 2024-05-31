package label

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_label resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_label", func(r *config.Resource) {
		r.ShortGroup = "label"
		r.References["project_id"] = config.Reference{
			Type: "github.com/globallogicuki/provider-harbor/apis/project/v1alpha1.Project",
		}
	})
}
