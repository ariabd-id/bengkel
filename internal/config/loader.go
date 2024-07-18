package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when load env: %s", err.Error())
	}

	return Config{
		Srv: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
