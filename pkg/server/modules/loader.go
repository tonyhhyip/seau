package modules

import (
	"fmt"
	"plugin"

	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/config"
)

func NewLoaderWithConfig(config *config.Config) *NativePluginLoader {
	return NewLoader(config.Plugin.RootPath)
}

func NewLoader(rootPath string) *NativePluginLoader {
	return &NativePluginLoader{
		RootPath: rootPath,
	}
}

type NativePluginLoader struct {
	RootPath string
}

func (n *NativePluginLoader) Load(name string) (p api.Plugin, err error) {
	filename := fmt.Sprintf("%s/%s_plugin.so", n.RootPath, name)
	ext, err := plugin.Open(filename)
	if err != nil {
		return
	}

	sym, err := ext.Lookup(api.PluginSymbol)
	if err != nil {
		return
	}

	p = sym.(api.Plugin)
	return
}
