package quickstart

import (
	"database/sql"
	"path/filepath"

	"github.com/murfffi/getaduck/download"
	"github.com/sclgo/adbcduck-go"
)

func init() {
	libFile, err := download.Do(download.DefaultSpec())
	if err != nil {
		panic(err)
	}
	absFile, err := filepath.Abs(libFile)
	if err != nil {
		panic(err)
	}
	sql.Register(adbcduck.DriverName, adbcduck.Make(absFile))
}
