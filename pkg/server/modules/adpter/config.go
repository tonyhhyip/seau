package adpter

import "github.com/tonyhhyip/seau/api"

type config struct {
	opener         api.Opener
	domainRegistry api.DomainRegistry
	blobManager    api.BlobManager
}

func (c *config) Opener() api.Opener {
	return c.opener
}

func (c *config) DomainRegistry() api.DomainRegistry {
	return c.domainRegistry
}

func (c *config) BlobManager() api.BlobManager {
	return c.blobManager
}
