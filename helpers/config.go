package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig() (string, error) {
	var err error
	if err = godotenv.Load("../.env"); err != nil {
		return "", err
	}

	return os.Getenv("API_KEY"), nil
}
