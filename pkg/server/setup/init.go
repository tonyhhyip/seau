package setup

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/tonyhhyip/seau/api"
	serverConfig "github.com/tonyhhyip/seau/pkg/server/config"
	"github.com/tonyhhyip/seau/pkg/server/db"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	pluginConfig "github.com/tonyhhyip/seau/pkg/server/modules/config"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

var handler http.Handler
var registry *modules.Registry

func init() {
	config := serverConfig.BuildFromEnv()
	opener := db.NewFromConfig(config)
	store := newStore(opener)
	initRegistry(config, store, opener)
	handler = newHandler(store, registry)
	initPlugin(config)
}

func initRegistry(config *serverConfig.Config, store repository.Store, opener api.Opener) {
	factory := &pluginConfig.DomainRegistryFactory{
		Store: store,
	}
	configFactory := newConfigFactory(opener, factory)
	loader := modules.NewLoaderWithConfig(config)
	registry = newRegistry(loader, configFactory)
}

func initPlugin(config *serverConfig.Config) {
	for _, name := range config.Plugin.Plugins {
		registry.Register(name)
	}
}
