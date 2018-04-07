package setup

import (
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

func newStore(opener api.Opener) *repository.Store {
	store := &repository.Store{
		Opener: opener,
	}
	return store
}
