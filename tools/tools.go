package tools

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %s", err.Error())
	}
	return nil
}
