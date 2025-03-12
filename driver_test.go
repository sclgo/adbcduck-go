package adbcduck_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sclgo/adbcduck-go"
	_ "github.com/sclgo/adbcduck-go/quickstart"
)

func TestE2E(t *testing.T) {
	db, err := sql.Open(adbcduck.DriverName, "")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = db.Close()
	})
	err = db.Ping()
	require.NoError(t, err)
	tx, err := db.Begin()
	require.NoError(t, err)
	require.NoError(t, tx.Commit())
	err = db.Close()
	require.NoError(t, err)
}
