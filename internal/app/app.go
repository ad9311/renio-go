package app

import (
	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/ad9311/renio-go/internal/model"
)

func Init() error {
	if err := conf.Init(); err != nil {
		return err
	}

	if err := db.Init(); err != nil {
		return err
	}

	if conf.GetEnv().AppEnv == conf.Production {
		if err := db.Migrate(); err != nil {
			return err
		}
	}

	if conf.GetEnv().Seed {
		if err := seed.Run(); err != nil {
			return err
		}
	}

	model.RegisterModels()

	return nil
}
