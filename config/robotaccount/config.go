package robotaccount

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_robot_account resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_robot_account", func(r *config.Resource) {
		r.ShortGroup = "robotaccount"
		r.Kind = "RobotAccount"
	})
}
