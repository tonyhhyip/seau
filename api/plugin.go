package api

import (
	"context"
	"net/http"
)

const PluginSymbol = "Plugin"

type Plugin interface {
	// Get unique ID of plugin
	ID() string
	// Get display name of plugin
	Name() string
	// Get handler for request
	Handler() http.Handler
	// Initial Setup
	Init(context.Context) error
	// SetConfig
	SetConfig(Config)
}
