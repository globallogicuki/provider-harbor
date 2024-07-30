package project

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/globallogicuki/provider-harbor/config/common"
)

// Configure harbor_project resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
	p.AddResourceConfigurator("harbor_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Project"
		r.References["registry_id"] = config.Reference{
			Type:      "github.com/globallogicuki/provider-harbor/apis/registry/v1alpha1.Registry",
			Extractor: common.RegistryIDExtractor,
		}
	})
}
