package envs

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	ENV         string
	PORT        string
	DatabaseURL string
	Seed        bool
	JWTToken    string
}

const (
	Development = "development"
	Test        = "test"
	Production  = "production"
	defaultPort = "8080"
	defaultENV  = Development
)

var (
	envVars   EnvVars
	once      sync.Once
	validENVs = map[string]bool{
		Development: true,
		Test:        true,
		Production:  true,
	}
)

func Init(rootDir string) error {
	var envErr error

	once.Do(func() {
		env := os.Getenv("APP_ENV")

		if !isValidENV(env) {
			envErr = fmt.Errorf("%s is not a valid environment", env)
			return
		}

		if env != Production {
			dotfile := fmt.Sprintf("%s/.env", rootDir)
			if err := godotenv.Load(dotfile); err != nil {
				envErr = fmt.Errorf("dofile: %s", err.Error())
				return
			}
		}

		envVars.setENV(env)
		envVars.setDatabaseURL(env)
		envVars.setSeeds()
		envVars.setPORT()
		if err := envVars.setJWTToken(); err != nil {
			envErr = err
		}
	})

	return envErr
}

func GetEnvs() EnvVars {
	return envVars
}

// --- Helpers --- //

func isValidENV(env string) bool {
	if env == "" {
		return true
	}

	return validENVs[env]
}

func (e *EnvVars) setENV(env string) {
	if env == "" {
		e.ENV = defaultENV
		return
	}
	e.ENV = env
}

func (e *EnvVars) setDatabaseURL(env string) {
	if env == Test {
		e.DatabaseURL = os.Getenv("TEST_DATABASE_URL")
	} else {
		e.DatabaseURL = os.Getenv("DATABASE_URL")
	}
}

func (e *EnvVars) setSeeds() {
	if os.Getenv("SEED") == "on" {
		e.Seed = true
		return
	}
	e.Seed = false
}

func (e *EnvVars) setPORT() {
	port := os.Getenv("PORT")
	if port == "" {
		e.PORT = defaultPort
		return
	}
	e.PORT = port
}

func (e *EnvVars) setJWTToken() error {
	if os.Getenv("JWT_TOKEN") == "" {
		return fmt.Errorf("JWT_TOKEN is not set")
	}
	e.JWTToken = os.Getenv("JWT_TOKEN")

	return nil
}
