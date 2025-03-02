package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App  AppConfig
	Auth AuthConfig
	Db   DbConfig
}

type AppConfig struct {
	Port int
}

type AuthConfig struct {
	Secret string
}

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file. Using default values")
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	return &Config{
		App: AppConfig{
			Port: port,
		},
		Auth: AuthConfig{
			Secret: os.Getenv("AUTH_SECRET"),
		},
		Db: DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Dbname:   os.Getenv("DB_NAME"),
			Sslmode:  os.Getenv("DB_SSLMODE"),
		},
	}
}
