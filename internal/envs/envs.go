package envs

import (
	"fmt"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/joho/godotenv"
)

var dotfile string

func Init() error {
	env := os.Getenv("ENV")

	switch env {
	case "development":
		dotfile = ".env"
	case "production":
		dotfile = ".env.production"
	case "test":
		dotfile = ".env.test"
	default:
		if err := os.Setenv("ENV", "development"); err != nil {
			return err
		}
		env = os.Getenv("ENV")
		dotfile = ".env"
	}

	if err := godotenv.Load(dotfile); err != nil {
		return err
	}
	console.Success(fmt.Sprintf("Loaded .env file in %s environment", env))

	return nil
}
