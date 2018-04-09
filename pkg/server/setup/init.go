package setup

import (
	"sync"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/seau/pkg/server/blob"
	serverConfig "github.com/tonyhhyip/seau/pkg/server/config"
	"github.com/tonyhhyip/seau/pkg/server/db"
)

var globalConfig *serverConfig.Config

func Bootstrap() {
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
	logrus.Debug("Load Plugin")
	group := new(sync.WaitGroup)
	loadPlugin := func(name string) {
		group.Add(1)
		logrus.Debugf("Plugin: %s", name)
		if err := registry.Register(name); err != nil {
			logrus.Error(err)
			panic(err)
		}
		group.Done()
	}
	for _, name := range globalConfig.Plugin.Plugins {
		go loadPlugin(name)
	}
	group.Wait()
}
