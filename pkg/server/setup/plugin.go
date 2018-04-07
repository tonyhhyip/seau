package setup

import (
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/modules/config"
)

func newConfigFactory(opener api.Opener, factory *config.DomainRegistryFactory) *config.PluginConfigFactory {
	return &config.PluginConfigFactory{
		Opener:                opener,
		DomainRegistryFactory: factory,
	}
}
