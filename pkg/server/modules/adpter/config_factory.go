package adpter

import "github.com/tonyhhyip/seau/api"

type PluginConfigFactory struct {
	Opener                api.Opener
	DomainRegistryFactory *DomainRegistryFactory
	BlobManagerFactory    *BlobFactory
}

func (pcf *PluginConfigFactory) Create(name string) *config {
	return &config{
		opener:         pcf.Opener,
		domainRegistry: pcf.DomainRegistryFactory.Create(name),
		blobManager:    pcf.BlobManagerFactory.Create(name),
	}
}
