package replication

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_replication resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_replication", func(r *config.Resource) {
		r.ShortGroup = "registry"
		r.Kind = "Replication"
		r.References["registry_id"] = config.Reference{
			Type: "github.com/globallogicuki/provider-harbor/apis/registry/v1alpha1.Registry",
		}
	})
}
