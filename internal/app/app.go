package app

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetRootDir() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to determine the root directory")
	}

	root := filepath.Join(filepath.Dir(filename), "../..")
	root, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("failed to resolve root path, %v", err)
	}

	return root, nil
}
