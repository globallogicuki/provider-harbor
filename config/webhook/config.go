package webhook

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_project_webhook resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_project_webook", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Webhook"
		r.References["project_id"] = config.Reference{
			Type: "github.com/globallogicuki/provider-harbor/apis/project/v1alpha1.Project",
		}
	})
}
