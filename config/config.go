package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	return os.Getenv(key)
}
