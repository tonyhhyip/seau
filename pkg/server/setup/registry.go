package setup

import (
	"sync"

	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/modules/config"
)

func newRegistry(loader modules.Loader, configFactory *config.PluginConfigFactory) *modules.Registry {
	return &modules.Registry{
		Loader:        loader,
		RegisterTable: new(sync.Map),
		ConfigFactory: configFactory,
	}
}
