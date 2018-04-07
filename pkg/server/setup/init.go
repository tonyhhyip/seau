package setup

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/tonyhhyip/seau/pkg/server/blob"
	serverConfig "github.com/tonyhhyip/seau/pkg/server/config"
	"github.com/tonyhhyip/seau/pkg/server/db"
)

var globalConfig *serverConfig.Config

func init() {
	globalConfig = serverConfig.BuildFromEnv()
	opener := db.NewFromConfig(globalConfig)
	blobManager, err := blob.NewFromConfig(globalConfig)
	if err != nil {
		panic(err)
	}
	store := newStore(opener)
	initRegistry(store, opener, blobManager)
	handler = newHandler(store, registry)
	initPlugin()
}

func initPlugin() {
	for _, name := range globalConfig.Plugin.Plugins {
		registry.Register(name)
	}
}
