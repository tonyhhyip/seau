package plugin

import (
	"context"
	"net/http"

	"github.com/tonyhhyip/seau/api"
)

type Plugin struct {
	processor *processContainer
	handler   http.Handler
}

func (p *Plugin) ID() string {
	return "php"
}

func (p *Plugin) Name() string {
	return "PHP Composer Repository"
}

func (p *Plugin) Init(context.Context) error {
	p.handler = p.processor.createHandler()
	return nil
}

func (p *Plugin) SetConfig(config api.Config) {
	p.processor.opener = config.Opener()
	p.processor.registry = config.DomainRegistry()
}

func (p *Plugin) Handler() http.Handler {
	return p.handler
}
