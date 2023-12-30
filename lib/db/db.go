package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/a10adotapp/a10a.app/ent"
	_ "github.com/a10adotapp/a10a.app/ent/runtime"
	"github.com/go-sql-driver/mysql"
)

func NewEntClient() *ent.Client {
	options := make([]ent.Option, 0)

	if os.Getenv("DEBUG_SQL") == "true" {
		options = append(options, ent.Debug())
	}

	client, err := ent.Open("mysql", newDSN(), options...)
	if err != nil {
		slog.Error(
			"NewEntClient",
			slog.Any("error", err),
		)
	}

	return client
}

// https://entgo.io/docs/transactions/#best-practices
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func newDSN() string {
	mysqlConfig := mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		DBName:               os.Getenv("DB_NAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	return mysqlConfig.FormatDSN()
}
