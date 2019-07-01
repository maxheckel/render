package server

import "github.com/kelseyhightower/envconfig"

type Config struct{
	DBHost     string `envconfig:"DB_HOST" default:"db"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"vinyl-analytics"`
	DBUser     string `envconfig:"DB_USER" default:"vinyl"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"secret"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}


func (a *App) BuildConfig() {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	a.Config = *config
}


func Load() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("", c)
	return c, err
}