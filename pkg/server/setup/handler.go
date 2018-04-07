package setup

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/tonyhhyip/seau/pkg/server"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

func newHandler(store repository.Store, registry *modules.Registry) http.Handler {
	handler := &server.Handler{
		Store:    store,
		Registry: registry,
	}

	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}

func GetHandler() http.Handler {
	return handler
}
