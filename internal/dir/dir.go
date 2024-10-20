package dir

import (
	"os"
	"path/filepath"
)

func RootDir() (string, error) {
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

func MigrationsDir() (string, error) {
	rootDir, err := RootDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(rootDir, "./db/migrations"), nil
}
