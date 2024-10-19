package db

import (
	"context"
	"fmt"
	"log"
	"os"
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

func QueryRow(dbExec QueryExe) error {
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

func (x *QueryExe) Query() error {
	ctx := context.Background()
	pool := GetPool()

	// Execute the query
	rows, err := pool.Query(ctx, x.QueryStr, x.QueryArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		// Create a new instance of the model for each row
		newModelPtr := reflect.New(reflect.TypeOf(x.Model)).Interface()

		// Prepare ScanArgs for the new model instance
		x.ScanArgs = spreadValues(newModelPtr)

		// Scan the row into the new model instance
		if err := rows.Scan(x.ScanArgs...); err != nil {
			return err
		}

		// Append the new model to the ModelSlice
		*x.ModelSlice = append(*x.ModelSlice, newModelPtr)
	}

	return nil
}

// --- Helpers --- //

func spreadValues(model any) []any {
	v := reflect.ValueOf(model).Elem()

	// Slice to hold pointers to the fields of the struct
	var values []any

	// Loop over the fields and append their addresses
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanAddr() {
			values = append(values, field.Addr().Interface())
		}
	}

	return values
}
