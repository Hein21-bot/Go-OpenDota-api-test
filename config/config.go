package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using system env")
	}
	fmt.Println("No error loading .env file")

}

func GetEnv(key string) string {
	return os.Getenv(key)
}
