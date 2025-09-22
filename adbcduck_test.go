package adbcduck_test

import (
	"database/sql"
	"fmt"
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

	logDuckdbVersion(t, db)

	t.Run("ping", func(t *testing.T) {
		err = db.Ping()
		require.NoError(t, err)
	})
	t.Run("transaction", func(t *testing.T) {
		tx, err := db.Begin()
		require.NoError(t, err)
		var version string
		err = db.QueryRow("SELECT VERSION()").Scan(&version)
		require.NoError(t, err)
		require.Equal(t, 2, strings.Count(version, "."))
		require.NoError(t, tx.Commit())
	})
	t.Run("ddl dml", func(t *testing.T) {
		require.NoError(t, exec(db, "create table foobar(n int)"))
		require.NoError(t, exec(db, "insert into foobar(n) values (?)", int8(1)))
		require.NoError(t, exec(db, "drop table foobar"))
	})
	t.Run("union", func(t *testing.T) {
		testUnion(t, db)
	})
	t.Run("decimal", func(t *testing.T) {
		var res any
		err = db.QueryRow("SELECT 0.13-0.07").Scan(&res)
		require.NoError(t, err)
		require.Equal(t, "0.06", fmt.Sprint(res))
	})
	t.Run("float", func(t *testing.T) {
		var res any
		err = db.QueryRow("SELECT 0.13::FLOAT-0.07").Scan(&res)
		require.NoError(t, err)
		require.Equal(t, "0.059999995", fmt.Sprint(res))
	})
	t.Run("s3", func(t *testing.T) {
		require.NoError(t, exec(db, "set s3_region='us-west-2'"))
		require.NoError(t, exec(db, "CREATE OR REPLACE VIEW earth AS SELECT * FROM read_parquet('s3://daylight-openstreetmap/earth/release=v1.58/*/*', hive_partitioning = true)"))
	})
	err = db.Close()
	require.NoError(t, err)
}

func logDuckdbVersion(t *testing.T, db *sql.DB) {
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	require.NoError(t, err)
	t.Log("duckdb version", version)
}

// Check for https://github.com/marcboeker/go-duckdb/issues/305
func testUnion(t *testing.T, db *sql.DB) {
	require.NoError(t, exec(db, "create table test(n int, a union(u varchar, v int))"))
	require.NoError(t, exec(db, "insert into test values(?, ?),(?, ?)", int8(1), "aba", int8(2), int8(2)))

	var uStr string
	var uInt int
	rows, err := db.Query("select a from test order by n")
	require.NoError(t, err)

	require.True(t, rows.Next(), "rows.Err() is %v", rows.Err())
	require.NoError(t, rows.Err())
	require.NoError(t, rows.Scan(&uStr))
	require.Equal(t, "aba", uStr)

	require.True(t, rows.Next(), "rows.Err() is %v", rows.Err())
	require.NoError(t, rows.Err())
	require.NoError(t, rows.Scan(&uInt))
	require.Equal(t, 2, uInt)
}

func exec(db *sql.DB, dml string, args ...any) error {
	_, err := db.Exec(dml, args...)
	return err
}
