package setup

import (
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/modules/config"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

func newConfigFactory(opener api.Opener, store *repository.Store) *config.PluginConfigFactory {
	return &config.PluginConfigFactory{
		Opener: opener,
		DomainRegistryFactory: &config.DomainRegistryFactory{
			Store: store,
		},
	}
}
