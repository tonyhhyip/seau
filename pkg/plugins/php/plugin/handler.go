package plugin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blang/semver"
	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/seau/api"
	jsonApi "github.com/tonyhhyip/seau/pkg/plugins/php/plugin/api"
	"github.com/tonyhhyip/seau/pkg/plugins/php/plugin/modals"
	"github.com/tonyhhyip/vodka"
)

const (
	notifyBatch = "/downloads"
	providerUrl = "/pkg/%package%/%hash%.json"
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

	router.GET("/packages", p.listProviders)
	router.GET("/p/:provider/:hash", p.listPackages)
	router.GET("/pkg/:provider/:package", p.listVersion)

	return router.Handler()
}

func (p *processContainer) listProviders(c *vodka.Context) {
	var err error
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

func (p *processContainer) listPackages(c *vodka.Context) {
	var err error
	provider, err := vodka.String(c.UserValue("provider"))
	if err != nil {
		c.Logger().Error(err.Error())
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	packages, err := p.store.GetPackagesByVendor(c.Host(), provider)
	if err != nil {
		c.Logger().Error(err.Error())
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	providers := make(map[string]jsonApi.Hash)
	for _, pkg := range packages {
		providers[provider+"/"+pkg.Name] = jsonApi.Hash{
			Sha256: pkg.Hash,
		}
	}

	c.JSON(http.StatusOK, jsonApi.Vendor{
		Providers: providers,
	})
}

func (p *processContainer) listVersion(c *vodka.Context) {
	var err error
	provider, err := vodka.String(c.UserValue("provider"))
	if err != nil {
		c.Logger().Error(err.Error())
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	pkg, err := vodka.String(c.UserValue("package"))
	if err != nil {
		c.Logger().Error(err.Error())
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	hostname := c.Host()

	versions, err := p.store.GetPackageVersion(hostname, provider, pkg)
	if err != nil {
		c.Logger().Error(err.Error())
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	pkgs := make(map[string]jsonApi.PackageVersion)

	for _, version := range versions {
		var pv jsonApi.PackageVersion
		if err := json.Unmarshal([]byte(version.ComposerContent), &pv); err != nil {
			c.Logger().Error(err.Error())
			c.Error(err.Error(), http.StatusInternalServerError)
			return
		}
		pv.Version = version.Version
		pv.Dist.Type = "tar"
		pv.Dist.ShaSum = version.ShaSum
		pv.Dist.Url = fmt.Sprintf("https://%s/b/%s/%s/%s", hostname, provider, pkg, version.Version)
		pv.Time = version.ReleaseTime.UTC().Format("2006-01-02T15:05:05-07:00")
		if version.Version == "dev-master" {
			pv.VersionNormalized = "9999999-dev"
		} else {
			normalizedVersion, _ := semver.Make(version.Version)
			pv.VersionNormalized = normalizedVersion.String() + ".0"
		}
		pkgs[version.Version] = pv
	}

	c.JSON(http.StatusOK, jsonApi.Packages{
		Packages: map[string]jsonApi.Package{
			provider + "/" + pkg: pkgs,
		},
	})
}
