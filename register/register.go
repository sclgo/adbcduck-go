// Package register adds a new driver instance with default library name to the database/sql driver set.
// Unlike package quickstart, it does not automatically download duckdb lib.
package register

import (
	"database/sql"

	"github.com/sclgo/adbcduck-go"
)

func init() {
	sql.Register(adbcduck.DriverName, adbcduck.Make("duckdb"))
}
