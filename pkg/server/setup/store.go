package setup

import (
	"github.com/dgrijalva/lfu-go"
	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

func newStore(opener api.Opener) repository.Store {
	cache := lfu.New()
	cache.LowerBound = 2
	cache.UpperBound = 16

	return &repository.CachedStore{
		Store: &repository.PostgresStore{
			Opener: opener,
		},
		Cache: cache,
	}
}
