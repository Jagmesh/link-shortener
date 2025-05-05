package config

import (
	"link-shortener/pkg/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var log = logger.GetWithScopes("MAIN")

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

type DbConfigCredentials struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Sslmode  string
}
type DbConfig struct {
	Credentials      DbConfigCredentials
	MaxRetriesNumber uint8
}

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file. Using default ENV values")
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}
	maxRetriesNumber, err := strconv.Atoi(os.Getenv("DB_MAX_RETRIES"))
	if err != nil {
		maxRetriesNumber = 3
	}

	return &Config{
		App: AppConfig{
			Port: port,
		},
		Auth: AuthConfig{
			Secret: os.Getenv("AUTH_SECRET"),
		},
		Db: DbConfig{
			Credentials: DbConfigCredentials{
				Host:     os.Getenv("DB_HOST"),
				Port:     dbPort,
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				Dbname:   os.Getenv("DB_NAME"),
				Sslmode:  os.Getenv("DB_SSLMODE"),
			},
			MaxRetriesNumber: uint8(maxRetriesNumber),
		},
	}
}
