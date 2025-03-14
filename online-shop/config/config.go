package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBSSLMode         string
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBMaxConnLifeTime int
}

func LoadConfig() *Config {
	err := godotenv.Load("environment/.env")
	if err != nil {
		log.Fatal("Error while fetching .env file for database")
	}

	MaxOpenConns, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_OPEN_CONNS"))
	MaxIdleConns, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNS"))
	MaxConnLifeTime, _ := strconv.Atoi(os.Getenv("POSTGRES_CONN_MAX_LIFETIME"))

	return &Config{
		DBHost:            os.Getenv("POSTGRES_HOST"),
		DBPort:            os.Getenv("POSTGRES_PORT"),
		DBUser:            os.Getenv("POSTGRES_USER"),
		DBPassword:        os.Getenv("POSTGRES_PASSWORD"),
		DBName:            os.Getenv("POSTGRES_NAME"),
		DBSSLMode:         os.Getenv("POSTGRES_SSLMODE"),
		DBMaxOpenConns:    MaxOpenConns,
		DBMaxIdleConns:    MaxIdleConns,
		DBMaxConnLifeTime: MaxConnLifeTime,
	}

}
