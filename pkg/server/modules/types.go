package modules

import "github.com/tonyhhyip/seau/api"

type Loader interface {
	Load(name string) (api.Plugin, error)
}
