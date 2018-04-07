package modules

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/seau/pkg/server/modules/adpter"
)

type Registry struct {
	Loader        Loader
	RegisterTable *sync.Map
	ConfigFactory *adpter.PluginConfigFactory
}

func (r *Registry) Register(name string) (err error) {
	_, exists := r.RegisterTable.Load(name)
	if exists {
		return errors.New("plugin already register")
	}

	plugin, err := r.Loader.Load(name)
	if err != nil || plugin == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = plugin.Init(ctx)
	if err != nil {
		return
	}

	config := r.ConfigFactory.Create(plugin.ID())
	plugin.SetConfig(config)

	r.RegisterTable.Store(name, plugin)

	logrus.Infof("Load plugin %s(%s)", plugin.Name(), plugin.ID())

	return
}
