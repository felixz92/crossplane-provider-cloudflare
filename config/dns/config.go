package dns

import "github.com/crossplane/upjet/pkg/config"

// ZoneRef references a zone by uuid
var ZoneRef = config.Reference{
	TerraformName: "cloudflare_zone",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "dns"
	})
	p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {
		r.ShortGroup = "dns"

		r.References["zone_id"] = ZoneRef
	})
}
