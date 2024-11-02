package conf

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"text/template"
)

var once sync.Once

func Init() error {
	var err error

	once.Do(func() {
		if err = InitEnv(); err != nil {
			return
		}

		if GetEnv().AppEnv != Test {
			InitSessionManager()

			_, err = BuildTemplateCache(template.FuncMap{})
			if err != nil {
				return
			}

		}
	})

	return err
}

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
