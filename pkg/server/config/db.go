package config

import "os"

type Database struct {
	Driver     string
	DataSource string
}

func dbFromEnv() *Database {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}
	return &Database{
		Driver:     driver,
		DataSource: os.Getenv("DB_URL"),
	}
}
