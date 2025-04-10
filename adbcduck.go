// Package adbcduck defines a DuckDB database/sql driver over the Arrow ADBC API
package adbcduck // import "github.com/sclgo/adbcduck-go"

import (
	"database/sql/driver"

	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	"github.com/sclgo/adbcduck-go/internal/sqldriver"
)

// DriverName defines the common driver name which packages ./quickstart and ./register
// use for sql.Register. Note that this package does not call sql.Register and does not have
// init() at all, so it can be imported without side-effects.
const DriverName = "adbcduck"

// Make creates database/sql/driver.Driver implementation for DuckDB over ADBC
//
// libraryName is the location of the duckdb shared library to use, either:
//   - full path to the library file (recommended)
//   - a bare file name like "libduckdb.so" which will be expected in the system library directories e.g. /lib
//     or on a directory included in LD_LIBRARY_PATH variable or an equivalent variable depending on the platform.
//     On Windows only, the library can be loaded from the working directory as well.
//   - a name like "duckdb" which will be automatically converted to a platform-specific bare file name,
//     as described in https://arrow.apache.org/adbc/main/cpp/driver_manager.html#usage .
//     The resulting bare file name is handled same as above.
//
// See ./register and ./quickstart packages for examples. In particular, quickstart demonstrates
// how to automatically download the library.
func Make(libraryName string) driver.Driver {
	return &sqldriver.Driver{
		Driver:      drivermgr.Driver{},
		LibraryName: libraryName,
	}
}
