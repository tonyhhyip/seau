package setup

import (
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/modules/adpter"
)

func newConfigFactory(
	opener api.Opener,
	domain *adpter.DomainRegistryFactory,
	blob *adpter.BlobFactory,
) *adpter.PluginConfigFactory {
	return &adpter.PluginConfigFactory{
		Opener:                opener,
		DomainRegistryFactory: domain,
		BlobManagerFactory:    blob,
	}
}
