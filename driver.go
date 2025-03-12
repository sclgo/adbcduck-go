// Package adbcduck defines a DuckDB database/sql driver over the Arrow ADBC API
package adbcduck // import "github.com/sclgo/adbcduck-go"

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	"github.com/apache/arrow-adbc/go/adbc/sqldriver"
)

const DriverName = "adbcduck"

// Make creates a new Driver instance, implementing database/sql/driver.Driver
//
// libraryName is the location of the duckdb shared library to use, either:
//   - full path to the library file (recommended)
//   - a name like "duckdb" which will be automatically converted to a platform-specific bare file name,
//     as described in https://arrow.apache.org/adbc/main/cpp/driver_manager.html#usage
//   - a bare file name like "libduckdb.so" which will be expected in the system library directories e.g. /lib
//     or on a directory included in LD_LIBRARY_PATH variable or an equivalent variable depending on the platform.
//     Only on Windows, that includes the working directory.
//
// See register and quickstart packages for examples.
func Make(libraryName string) *Driver {
	return &Driver{
		adbcDriver: sqldriver.Driver{
			Driver: drivermgr.Driver{},
		},
		libraryName: libraryName,
	}
}

const namePattern = "driver=%s;entrypoint=duckdb_adbc_init;path=%s"

// Driver implements a DuckDB database/sql driver over the Arrow ADBC API. Use Make to create.
type Driver struct {
	adbcDriver  sqldriver.Driver
	libraryName string
}

// Interface validation for Driver
var _ driver.Driver = (*Driver)(nil)
var _ driver.DriverContext = (*Driver)(nil)

type duckDbConnector struct {
	driver.Connector
}

// Connect implements database/sql/driver.Connector
func (d duckDbConnector) Connect(ctx context.Context) (driver.Conn, error) {
	adbcConn, err := d.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &conn{
		fullConn: adbcConn.(fullConn),
	}, nil
}

// OpenConnector implements database/sql/driver.DriverContext
func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	adbcCtr, err := d.adbcDriver.OpenConnector(d.getDsn(name))
	return duckDbConnector{adbcCtr}, err
}

// Open implements database/sql/driver.Driver
func (d *Driver) Open(name string) (driver.Conn, error) {
	connector, err := d.OpenConnector(name)
	if err != nil {
		return nil, err
	}
	return connector.Connect(context.Background())
}

func (d *Driver) getDsn(name string) string {
	return fmt.Sprintf(namePattern, d.libraryName, name)
}
