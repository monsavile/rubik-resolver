package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const envFileName = ".env"

func Load() error {
	if _, err := os.Stat(envFileName); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if err := godotenv.Load(envFileName); err != nil {
		return err
	}

	return nil
}
