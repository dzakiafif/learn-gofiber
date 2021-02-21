package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetConfigEnv(key string) string {
	err := godotenv.Load()
	PanicHandler(err)

	return os.Getenv(key)
}