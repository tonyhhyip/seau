package plugin

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/seau/api"
	jsonApi "github.com/tonyhhyip/seau/pkg/plugins/php/plugin/api"
	"github.com/tonyhhyip/seau/pkg/plugins/php/plugin/modals"
	"github.com/tonyhhyip/vodka"
)

const (
	notifyBatch = "/downloads"
	providerUrl = "/p/%package%/%hash%.json"
)

type processContainer struct {
	opener   api.Opener
	registry api.DomainRegistry
	blob     api.BlobManager
	store    *modals.Store
}

func (p *processContainer) init(config api.Config) {
	p.opener = config.Opener()
	p.registry = config.DomainRegistry()
	p.blob = config.BlobManager()

	p.store = &modals.Store{
		Opener: p.opener,
	}
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
	router.Use(vodka.MiddlewareFunc(stripJson))
	router.Use(vodka.MiddlewareFunc(func(next vodka.Handler) vodka.Handler {
		return vodka.HandlerFunc(func(c *vodka.Context) {
			c.Logger().Infof("%s %s %s", c.Request.Method, c.Host(), c.URL())
			next.Handle(c)
		})
	}))

	router.GET("/packages", p.listPackages)

	return router.Handler()
}

func (p *processContainer) listPackages(c *vodka.Context) {
	vendors, err := p.store.GetVendorByDomain(c.Host())
	if err != nil {
		c.Logger().Error(err.Error())
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	includes := make(map[string]jsonApi.Hash)
	for _, vendor := range vendors {
		includes["/p/"+vendor.Name+"/%hash%.json"] = jsonApi.Hash{
			Sha256: vendor.Hash,
		}
	}

	list := jsonApi.ProviderList{
		Packages:         []string{},
		Includes:         make(map[string]jsonApi.Hash),
		NotifyBatch:      notifyBatch,
		ProvidersUrl:     providerUrl,
		ProviderIncludes: includes,
	}

	c.JSON(http.StatusOK, list)
}
