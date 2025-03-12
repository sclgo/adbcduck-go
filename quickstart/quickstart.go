// Package quickstart downloads DuckDB library from Github and registers a new driver instance with it
package quickstart

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/murfffi/getaduck/download"
	"github.com/sclgo/adbcduck-go"
)

func init() {
	spec := download.DefaultSpec()
	if ver := os.Getenv("DUCKDB_VERSION"); ver != "" {
		spec.Version = ver
	}
	libFile, err := download.Do(spec)
	if err != nil {
		panic(err)
	}
	absFile, err := filepath.Abs(libFile)
	if err != nil {
		panic(err)
	}
	sql.Register(adbcduck.DriverName, adbcduck.Make(absFile))
}
