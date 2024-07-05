package purgeauditlog

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_purge_audit_log resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_purge_audit_log", func(r *config.Resource) {
		r.ShortGroup = "purgeauditlog"
		r.Kind = "PurgeAuditLog"
	})
}
