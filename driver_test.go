package adbcduck_test

import (
	"database/sql"
	"strings"
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
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	require.NoError(t, err)
	require.Equal(t, 2, strings.Count(version, "."))
	require.NoError(t, tx.Commit())
	err = db.Close()
	require.NoError(t, err)
}
