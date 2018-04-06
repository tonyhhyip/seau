package config

import "github.com/tonyhhyip/seau/api"

type PluginConfigFactory struct {
	Opener                api.Opener
	DomainRegistryFactory *DomainRegistryFactory
}

func (pcf *PluginConfigFactory) Create(name string) *config {
	return &config{
		opener:         pcf.Opener,
		domainRegistry: pcf.DomainRegistryFactory.Create(name),
	}
}
