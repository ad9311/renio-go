package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func Init() {
	once.Do(func() {
		fmt.Println("! Connecting to database...")

		dbURL := "postgres://ad9311:@localhost:5432/renio_go_development"

		var err error
		pool, err = pgxpool.New(context.Background(), dbURL)
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
