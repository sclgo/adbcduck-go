package adbcduck // import "github.com/sclgo/duckdb-adbc-go"

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	"github.com/apache/arrow-adbc/go/adbc/sqldriver"
)

const DriverName = "adbcduck"

var LibraryName = "duckdb"

func init() {
	sql.Register(DriverName, makeDriver())
}

func makeDriver() driver.Driver {
	return duckdbDriver{
		adbcDriver: sqldriver.Driver{
			Driver: drivermgr.Driver{},
		},
	}
}

const namePattern = "driver=%s;entrypoint=duckdb_adbc_init;path=%s"

type duckdbDriver struct {
	adbcDriver sqldriver.Driver
}

func (d duckdbDriver) OpenConnector(name string) (driver.Connector, error) {
	return d.adbcDriver.OpenConnector(getDsn(name))
}

func (d duckdbDriver) Open(name string) (driver.Conn, error) {
	return d.adbcDriver.Open(getDsn(name))
}

func getDsn(name string) string {
	return fmt.Sprintf(namePattern, LibraryName, name)
}
