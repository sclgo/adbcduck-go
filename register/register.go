package register

import (
	"database/sql"

	"github.com/sclgo/adbcduck-go"
)

func init() {
	sql.Register(adbcduck.DriverName, adbcduck.Make("duckdb"))
}
