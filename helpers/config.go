package helpers

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func LoadConfig() (string, int64, error) {
	var err error
	if err = godotenv.Load("./.env"); err != nil {
		return "", 0, err
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 0)

	if err != nil {
		port = 5000
	}

	return os.Getenv("API_KEY"), port, nil
}
