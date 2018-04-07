package config

import (
	"os"
	"strings"
)

type Plugin struct {
	RootPath string
	Plugins  []string
}

func pluginFromEnv() *Plugin {
	rootPath := os.Getenv("PLUGIN_PATH")
	if rootPath == "" {
		rootPath, _ = os.Getwd()
	}

	pluginNames := os.Getenv("PLUGINS")
	plugins := strings.Split(pluginNames, ",")

	return &Plugin{
		RootPath: rootPath,
		Plugins:  plugins,
	}
}
