package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server"
	"github.com/tonyhhyip/seau/pkg/server/db"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/modules/config"
	"github.com/tonyhhyip/seau/pkg/server/repository"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	opener := newOpener()
	store := newStore(opener)
	registry := newRegistry(opener, store)
	handler := &server.Handler{
		Store:    store,
		Registry: registry,
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, handlers.CombinedLoggingHandler(os.Stdout, handler))
}
func newOpener() api.Opener {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("DB_URL is not set")
	}
	return db.New("postgres", os.Getenv("DB_URL"))
}

func newStore(opener api.Opener) *repository.Store {
	store := &repository.Store{
		Opener: opener,
	}
	return store
}

func newRegistry(opener api.Opener, store *repository.Store) *modules.Registry {
	rootPath, _ := os.Getwd()
	loader := &modules.NativePluginLoader{
		RootPath: rootPath,
	}
	registry := &modules.Registry{
		Loader:        loader,
		RegisterTable: new(sync.Map),
		ConfigFactory: newConfigFactory(opener, store),
	}
	return registry
}

func newConfigFactory(opener api.Opener, store *repository.Store) *config.PluginConfigFactory {
	return &config.PluginConfigFactory{
		Opener: opener,
		DomainRegistryFactory: &config.DomainRegistryFactory{
			Store: store,
		},
	}
}
