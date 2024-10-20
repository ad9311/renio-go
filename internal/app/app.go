package app

import (
	"os"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/ad9311/renio-go/internal/envs"
)

func Init() error {
	if err := envs.Init(); err != nil {
		return err
	}

	if err := db.Init(); err != nil {
		return err
	}

	if os.Getenv("MIGRATE") == "on" {
		if err := migration.Up(); err != nil {
			return err
		}
	}

	if os.Getenv("SEED") == "on" {
		if err := seed.Run(); err != nil {
			return err
		}
	}

	return nil
}
