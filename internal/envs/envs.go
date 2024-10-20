package envs

import (
	"fmt"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/dir"
	"github.com/joho/godotenv"
)

var dotfile string

func Init() error {
	env := os.Getenv("ENV")

	rootDir, err := dir.RootDir()
	if err != nil {
		return err
	}

	switch env {
	case "development":
		dotfile = fmt.Sprintf("%s/.env", rootDir)
	case "production":
		dotfile = fmt.Sprintf("%s/.env.production", rootDir)
	case "test":
		dotfile = fmt.Sprintf("%s/.env.test", rootDir)
	default:
		if err := os.Setenv("ENV", "development"); err != nil {
			return err
		}
		env = os.Getenv("ENV")
		dotfile = fmt.Sprintf("%s/.env", rootDir)
	}

	if err := godotenv.Load(dotfile); err != nil {
		return err
	}
	console.Success(fmt.Sprintf("Loaded .env file in %s environment", env))

	return nil
}
