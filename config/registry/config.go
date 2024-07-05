package registry

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_registry resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_registry", func(r *config.Resource) {
		r.ShortGroup = "registry"
		r.Kind = "Registry"
	})
}
