package db

import (
	"database/sql"

	"github.com/tonyhhyip/seau/pkg/server/config"
)

func NewFromConfig(config *config.Config) *opener {
	return New(config.Database.Driver, config.Database.DataSource)
}

func New(driver, dataSource string) *opener {
	return &opener{
		Driver:     driver,
		DataSource: dataSource,
	}
}

type opener struct {
	Driver     string
	DataSource string
}

func (o *opener) Open() (*sql.DB, error) {
	return sql.Open(o.Driver, o.DataSource)
}
