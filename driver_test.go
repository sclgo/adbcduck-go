package duckadbc_test

import (
	"database/sql"
	"testing"

	_ "github.com/sclgo/duckdb-adbc-go"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	db, err := sql.Open("duckadbc", "")
	require.NoError(t, err)
	err = db.Close()
	require.NoError(t, err)
}
