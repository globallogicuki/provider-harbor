package tasks

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure harbor_tasks resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_tasks", func(r *config.Resource) {
		r.ShortGroup = "tasks"
		r.Kind = "Task"
	})
}
