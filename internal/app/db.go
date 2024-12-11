package app

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type QueryExe struct {
	QueryStr   string
	QueryArgs  []any
	ScanArgs   []any
	Model      any
	ModelSlice *[]any
}

var (
	pool  *pgxpool.Pool
	sqlDB *sql.DB
	once  sync.Once
)

func InitDB() error {
	var dbErr error

	once.Do(func() {
		var err error
		pool, err = pgxpool.New(context.Background(), GetEnv().DatabaseURL)
		if err != nil {
			dbErr = err
			return
		}

		if err = pool.Ping(context.Background()); err != nil {
			dbErr = err
			return
		}

		pgxConfig, parseErr := pgxpool.ParseConfig(GetEnv().DatabaseURL)
		if parseErr != nil {
			dbErr = parseErr
			return
		}

		sqlDB = stdlib.OpenDB(*pgxConfig.ConnConfig)
	})

	return dbErr
}

func GetPool() *pgxpool.Pool {
	return pool
}

func GetSQLDB() *sql.DB {
	return sqlDB
}

func (x *QueryExe) QueryRow() error {
	printQuery(x.QueryStr)

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
	printQuery(x.QueryStr)

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

func (x *QueryExe) Exec() error {
	printQuery(x.QueryStr)

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

func printQuery(query string) {
	if GetEnv().AppEnv != "test" {
		words := strings.Fields(query)
		query = strings.Join(words, " ")
		fmt.Printf("BEGIN `%s`\n", query)
	}
}
