package setup

import (
	"os"

	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/db"
)

func newOpener() api.Opener {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("DB_URL is not set")
	}
	return db.New("postgres", os.Getenv("DB_URL"))
}
