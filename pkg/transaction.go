package pkg

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Transaction - exec transaction
func Transaction(ctx context.Context, db *sqlx.DB, fn func(*sqlx.Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to BeginTx: %w", err)
	}

	succeed := false
	defer func() {
		if succeed {
			return
		}

		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to Rollback TX: %w", err)
		}
	}()

	if err = fn(tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		err = fmt.Errorf("failed to Commit TX: %w", err)
		return
	}

	succeed = true
	return nil
}
