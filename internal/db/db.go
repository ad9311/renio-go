package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func Init() {
	once.Do(func() {
		fmt.Println("! Connecting to database...")

		var err error
		pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatalf("x Unable to connect to database: %v\n", err)
		}

		if err = pool.Ping(context.Background()); err != nil {
			log.Fatalf("x Unable to ping database: %v\n", err)
		}

		fmt.Print("✓ Database connection established\n\n")
	})
}

func GetPool() *pgxpool.Pool {
	return pool
}

func Migrate() {
	db, err := goose.OpenDBWithDriver("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to open DB: %v\n", err)
	}
	defer db.Close()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to find current working directory: %v\n", err)
	}

	dir := filepath.Join(cwd, "./db/migrations")

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read migrations directory: %v\n", err)
	}

	if len(files) == 0 {
		fmt.Print("! Migrations directory is empty, skipping migrations\n\n")
		return
	}

	if err := goose.Up(db, dir); err != nil {
		log.Fatalf("failed to apply migrations: %v\n", err)
	}

	fmt.Println("")
}
