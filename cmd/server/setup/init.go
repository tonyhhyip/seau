package setup

import (
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/tonyhhyip/seau/pkg/server/modules"
)

var handler http.Handler
var registry *modules.Registry

func init() {
	opener := newOpener()
	store := newStore(opener)

	{
		path := os.Getenv("PLUGIN_PATH")
		if path == "" {
			path, _ = os.Getwd()
		}
		registry = newRegistry(path, opener, store)
	}
	handler = newHandler(store, registry)
	initPlugin()
}

func initPlugin() {
	pluginNames := os.Getenv("PLUGINS")
	plugins := strings.Split(pluginNames, ",")
	for _, name := range plugins {
		registry.Register(name)
	}
}
