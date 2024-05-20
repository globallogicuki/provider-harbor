package project

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
