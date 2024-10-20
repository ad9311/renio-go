package envs

import (
	"fmt"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/joho/godotenv"
)

var dotfile string

func Init() {
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
			console.Fatal(fmt.Sprintf("failed to set env, %s", err.Error()))
		}
		env = os.Getenv("ENV")
		dotfile = ".env"
	}

	err := godotenv.Load(dotfile)
	if err != nil {
		console.Fatal(fmt.Sprintf("could not load .env file, %s", err.Error()))
	}
	console.Success(fmt.Sprintf("Loaded .env file in %s environment", env))
}
