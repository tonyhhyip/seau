package config

type Config struct {
	Database *Database
	Plugin   *Plugin
	Blob     *Blob
}

func BuildFromEnv() *Config {
	return &Config{
		Database: dbFromEnv(),
		Plugin:   pluginFromEnv(),
		Blob:     blobFromEnv(),
	}
}
