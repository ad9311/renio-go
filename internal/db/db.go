package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBExec struct {
	QueryStr   string
	QueryArgs  []any
	ScanArgs   []any
	ModelSlice *[]any
}

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

func QueryRow(dbExec DBExec) error {
	ctx := context.Background()
	pool := GetPool()

	if err := pool.QueryRow(
		ctx,
		dbExec.QueryStr,
		dbExec.QueryArgs...,
	).Scan(
		dbExec.ScanArgs...,
	); err != nil {
		return err
	}

	return nil
}

func Query(dbExec DBExec) error {
	ctx := context.Background()
	pool := GetPool()

	rows, err := pool.Query(ctx, dbExec.QueryStr, dbExec.QueryArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var v []any
		if err := rows.Scan(dbExec.ScanArgs...); err != nil {
			return err
		}

		*dbExec.ModelSlice = append(*dbExec.ModelSlice, v)
	}

	return nil
}
