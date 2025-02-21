package adbcduck_test

import (
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/murfffi/getaduck/download"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sclgo/duckdb-adbc-go"
)

func TestE2E(t *testing.T) {
	libFile, err := download.Do(download.DefaultSpec())
	require.NoError(t, err)
	adbcduck.LibraryName, err = filepath.Abs(libFile)
	require.NoError(t, err)
	db, err := sql.Open(adbcduck.DriverName, "")
	require.NoError(t, err)
	err = db.Ping()
	assert.NoError(t, err)
	err = db.Close()
	require.NoError(t, err)
}
