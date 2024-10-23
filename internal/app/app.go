package app

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/ad9311/renio-go/internal/envs"
)

func Init() error {
	if err := envs.Init(); err != nil {
		return err
	}
	console.Success(fmt.Sprintf("Loading from %s environment", envs.GetEnvs().ENV))

	console.Info("Connecting to database...")
	if err := db.Init(); err != nil {
		return err
	}
	console.Success("Database connection established")

	if envs.GetEnvs().Seed {
		if err := seed.Run(); err != nil {
			return err
		}
	}

	return nil
}
