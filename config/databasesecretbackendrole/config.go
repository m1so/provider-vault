package DIR_databasesecretbackendrole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_database_secret_backend_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vault"
	})
}