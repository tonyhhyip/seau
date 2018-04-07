package plugin

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/vodka"
)

type processContainer struct {
	opener   api.Opener
	registry api.DomainRegistry
}

func (p *processContainer) createHandler() http.Handler {
	server := vodka.New("")
	logger := logrus.New()
	server.SetLogger(logger.WithField("source", "php"))
	server.StandBy(p.createRouter())
	return server.Server.Handler
}

func (p *processContainer) createRouter() vodka.Handler {
	router := vodka.NewRouter()
	router.Use(vodka.MiddlewareFunc(func(next vodka.Handler) vodka.Handler {
		return vodka.HandlerFunc(func(c *vodka.Context) {
			path := c.Request.URL.Path
			if strings.HasSuffix(path, ".json") {
				c.Request.URL.Path = path[:strings.LastIndex(path, ".json")]
			}
			next.Handle(c)
		})
	}))

	return router.Handler()
}
