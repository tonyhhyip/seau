package db

import "database/sql"

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
