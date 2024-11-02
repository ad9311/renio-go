package conf

import (
	"sync"
)

var once sync.Once

func Init() error {
	var err error

	once.Do(func() {
		if err = InitEnv(); err != nil {
			return
		}
	})

	return err
}
