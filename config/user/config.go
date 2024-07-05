package user

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_user resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_user", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "User"
	})
}
