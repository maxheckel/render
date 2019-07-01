package server

type Config struct{
	DBHost     string `envconfig:"DB_HOST" default:"db"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"vinyl-analytics"`
	DBUser     string `envconfig:"DB_USER" default:"vinyl"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"secret"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`

}