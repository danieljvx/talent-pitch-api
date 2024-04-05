package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print()
	}
	return os.Getenv(key)
}
