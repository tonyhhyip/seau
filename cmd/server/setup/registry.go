package setup

import (
	"sync"

	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

func newRegistry(path string, opener api.Opener, store *repository.Store) *modules.Registry {
	loader := &modules.NativePluginLoader{
		RootPath: path,
	}
	registry := &modules.Registry{
		Loader:        loader,
		RegisterTable: new(sync.Map),
		ConfigFactory: newConfigFactory(opener, store),
	}
	return registry
}
