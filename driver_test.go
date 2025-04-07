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

	testUnion(t, db)

	err = db.Close()
	require.NoError(t, err)
}

// Check for https://github.com/marcboeker/go-duckdb/issues/305
func testUnion(t *testing.T, db *sql.DB) {
	require.NoError(t, exec(db, "create table test(n int, a union(u varchar, v int))"))
	require.NoError(t, exec(db, "insert into test values(1, 'aba'),(2, 2)"))

	var uStr string
	var uInt int
	rows, err := db.Query("select a from test order by n")
	require.NoError(t, err)

	require.True(t, rows.Next())
	require.NoError(t, rows.Err())
	require.NoError(t, rows.Scan(&uStr))
	require.Equal(t, "aba", uStr)

	require.True(t, rows.Next())
	require.NoError(t, rows.Err())
	require.NoError(t, rows.Scan(&uInt))
	require.Equal(t, 2, uInt)
}

func exec(db *sql.DB, dml string) error {
	_, err := db.Exec(dml)
	return err
}
