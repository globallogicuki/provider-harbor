package project

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_project resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
