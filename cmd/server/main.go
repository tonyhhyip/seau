package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/tonyhhyip/seau/pkg/server"
	"github.com/tonyhhyip/seau/pkg/server/db"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/repository"

	_ "github.com/lib/pq"
)

func main() {
	store := newStore()
	registry := newRegistry()
	handler := &server.Handler{
		Store:    store,
		Registry: registry,
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, handler)
}

func newStore() *repository.Store {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("DB_URL is not set")
	}
	opener := db.New("postgres", os.Getenv("DB_URL"))
	store := &repository.Store{
		Opener: opener,
	}
	return store
}

func newRegistry() *modules.Registry {
	rootPath, _ := os.Getwd()
	loader := &modules.NativePluginLoader{
		RootPath: rootPath,
	}
	registry := &modules.Registry{
		Loader:        loader,
		RegisterTable: new(sync.Map),
	}
	return registry
}
