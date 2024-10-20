package db

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type QueryExe struct {
	QueryStr   string
	QueryArgs  []any
	ScanArgs   []any
	Model      any
	ModelSlice *[]any
}

var (
	pool *pgxpool.Pool
	once sync.Once
)

func Init(dabataseURL string) {
	once.Do(func() {
		fmt.Println("! Connecting to database...")

		var err error
		pool, err = pgxpool.New(context.Background(), dabataseURL)
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

func (x *QueryExe) QueryRow() error {
	fmt.Printf("BEGIN `%s`\n", x.QueryStr)
	ctx := context.Background()
	pool := GetPool()

	model := reflect.New(reflect.TypeOf(x.Model)).Interface()
	x.ScanArgs = spreadValues(model)

	if err := pool.QueryRow(
		ctx,
		x.QueryStr,
		x.QueryArgs...,
	).Scan(
		x.ScanArgs...,
	); err != nil {
		return err
	}

	x.Model = model

	return nil
}

func (x *QueryExe) Query() error {
	fmt.Printf("BEGIN `%s`\n", x.QueryStr)
	ctx := context.Background()
	pool := GetPool()

	rows, err := pool.Query(ctx, x.QueryStr, x.QueryArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		model := reflect.New(reflect.TypeOf(x.Model)).Interface()

		x.ScanArgs = spreadValues(model)
		if err := rows.Scan(x.ScanArgs...); err != nil {
			return err
		}

		*x.ModelSlice = append(*x.ModelSlice, model)
	}

	return nil
}

func (x *QueryExe) QueryWithoutScan() error {
	fmt.Printf("BEGIN `%s``\n", x.QueryStr)

	ctx := context.Background()
	pool := GetPool()

	if _, err := pool.Exec(ctx, x.QueryStr, x.QueryArgs...); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func spreadValues(model any) []any {
	v := reflect.ValueOf(model).Elem()

	var values []any

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanAddr() {
			values = append(values, field.Addr().Interface())
		}
	}

	return values
}
