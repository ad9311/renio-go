package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/ad9311/renio-go/internal/envs"
)

func Init() error {
	rootDir, err := getRootDir()
	if err != nil {
		return err
	}

	if err := envs.Init(rootDir); err != nil {
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

// --- Helpers --- //

func getRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", os.ErrNotExist
}
