package adbcduck

import (
	"context"
	"database/sql/driver"
	"errors"
)

type fullConn interface {
	driver.Conn
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.QueryerContext
}
type conn struct {
	fullConn

	tx bool
}

func (c *conn) BeginTx(ctx context.Context, _ driver.TxOptions) (driver.Tx, error) {
	if c.tx {
		return nil, errors.New("adbcduck: transaction already started")
	}
	err := c.exec(ctx, "BEGIN TRANSACTION")
	if err != nil {
		return nil, err
	}
	c.tx = true
	return tx{c}, nil
}

func (c *conn) exec(ctx context.Context, query string) error {
	stmt, err := c.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer func() {
		stmtErr := stmt.Close()
		if stmtErr != nil && err == nil {
			err = stmtErr
		}
	}()
	_, err = stmt.(driver.StmtExecContext).ExecContext(ctx, nil)
	return err
}

type tx struct {
	c *conn
}

func (t tx) Commit() error {
	err := t.c.exec(context.Background(), "COMMIT")
	// We assume that transaction is no longer active if commit failed.
	// Most likely reason for failure is that commit/rollback was manually called.
	t.c.tx = false
	return err
}

func (t tx) Rollback() error {
	err := t.c.exec(context.Background(), "ROLLBACK")
	t.c.tx = false // see comment above
	return err
}
