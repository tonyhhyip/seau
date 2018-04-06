package config

import "github.com/tonyhhyip/seau/api"

type config struct {
	opener         api.Opener
	domainRegistry api.DomainRegistry
}

func (c *config) Opener() api.Opener {
	return c.opener
}

func (c *config) DomainRegistry() api.DomainRegistry {
	return c.domainRegistry
}
