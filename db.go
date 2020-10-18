package pop

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func (db *DB) TransactionContext(ctx context.Context) (*Tx, error) {
	return newTX(ctx, db, nil)
}

func (db *DB) Transaction() (*Tx, error) {
	return newTX(context.Background(), db, nil)
}

func (db *DB) TransactionContextOptions(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	return newTX(ctx, db, opts)
}

func (db *DB) Rollback() error {
	return nil
}

func (db *DB) Commit() error {
	return nil
}
