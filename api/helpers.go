package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/elgris/sqrl"

	"github.com/johanbrandhorst/chronic-pain-tracker/api/internal"
)

const (
	// https://www.postgresql.org/docs/current/static/errcodes-appendix.html
	// This error code is seen sometimes when the database is starting.
	psqlCannotConnectNow = "57P03"
)

func ensureSchema(db *sql.DB, sb sqrl.StatementBuilderType) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// No-op if commit done first
	defer tx.Rollback()

	_, err = tx.Exec(fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s ("+
			"%s timestamptz NOT NULL,"+
			"%s text NOT NULL"+
			")",
		internal.EventsTable,
		internal.TimestampColumn,
		internal.PainLevelColumn,
	))
	if err != nil {
		return err
	}

	return tx.Commit()
}

// stdlibQueryRowContext is used to support `sqrl.QueryRowContextWith`
// with database/sql.DB
type stdlibQueryRowContext func(context.Context, string, ...interface{}) *sql.Row

func (s stdlibQueryRowContext) QueryRowContext(ctx context.Context, q string, args ...interface{}) sqrl.RowScanner {
	return s(ctx, q, args...)
}

// stdlibQueryRow is used to support `sqrl.QueryRowWith`
// with database/sql.DB
type stdlibQueryRow func(string, ...interface{}) *sql.Row

func (s stdlibQueryRow) QueryRow(q string, args ...interface{}) sqrl.RowScanner {
	return s(q, args...)
}
