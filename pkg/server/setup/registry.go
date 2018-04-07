package setup

import (
	"sync"

	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/blob"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/modules/adpter"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

var registry *modules.Registry

func initRegistry(store repository.Store, opener api.Opener, manager blob.Manager) {
	domainRegistryFactory := &adpter.DomainRegistryFactory{
		Store: store,
	}
	blobFactory := &adpter.BlobFactory{
		Manager: manager,
	}
	configFactory := newConfigFactory(opener, domainRegistryFactory, blobFactory)
	loader := modules.NewLoaderWithConfig(globalConfig)
	registry = newRegistry(loader, configFactory)
}

func newRegistry(loader modules.Loader, configFactory *adpter.PluginConfigFactory) *modules.Registry {
	return &modules.Registry{
		Loader:        loader,
		RegisterTable: new(sync.Map),
		ConfigFactory: configFactory,
	}
}
