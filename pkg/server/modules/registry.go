package modules

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Registry struct {
	Loader        Loader
	RegisterTable *sync.Map
}

func (r *Registry) Register(name string) (err error) {
	_, exists := r.RegisterTable.Load(name)
	if exists {
		return errors.New("plugin already register")
	}

	plugin, err := r.Loader.Load(name)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = plugin.Init(ctx)
	if err != nil {
		return
	}

	r.RegisterTable.Store(name, plugin)

	return
}
