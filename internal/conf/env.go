package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv      string
	Port        string
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
	env       *Env
	validENVs = map[string]bool{
		Development: true,
		Test:        true,
		Production:  true,
	}
)

func InitEnv() error {
	env = &Env{}

	appEnv := os.Getenv("APP_ENV")

	if !isValidENV(appEnv) {
		return fmt.Errorf("%s is not a valid environment", appEnv)
	}

	if appEnv != Production {
		rootDir, err := GetRootDir()
		if err != nil {
			return err
		}

		dotfile := fmt.Sprintf("%s/.env", rootDir)
		if err := godotenv.Load(dotfile); err != nil {
			return fmt.Errorf("dofile: %s", err.Error())
		}
	}

	env.setENV(appEnv)
	env.setDatabaseURL(appEnv)
	env.setSeeds()
	env.setPORT()
	if err := env.setJWTToken(); err != nil {
		return err
	}

	return nil
}

func GetEnv() *Env {
	return env
}

// --- Helpers --- //

func isValidENV(env string) bool {
	if env == "" {
		return true
	}

	return validENVs[env]
}

func (e *Env) setENV(env string) {
	if env == "" {
		e.AppEnv = defaultENV
		return
	}

	e.AppEnv = env
}

func (e *Env) setDatabaseURL(env string) {
	if env == Test {
		e.DatabaseURL = os.Getenv("TEST_DATABASE_URL")
	} else {
		e.DatabaseURL = os.Getenv("DATABASE_URL")
	}
}

func (e *Env) setSeeds() {
	if os.Getenv("SEED") == "on" {
		e.Seed = true
		return
	}
	e.Seed = false
}

func (e *Env) setPORT() {
	port := os.Getenv("PORT")
	if port == "" {
		e.Port = defaultPort
		return
	}
	e.Port = port
}

func (e *Env) setJWTToken() error {
	if os.Getenv("JWT_TOKEN") == "" {
		return fmt.Errorf("JWT_TOKEN is not set")
	}
	e.JWTToken = os.Getenv("JWT_TOKEN")

	return nil
}
