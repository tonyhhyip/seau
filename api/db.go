package api

import "database/sql"

type Opener interface {
	Open() (*sql.DB, error)
}
