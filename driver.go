package adbcduck // import "github.com/sclgo/adbcduck-go"

import (
	"database/sql/driver"
	"fmt"

	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	"github.com/apache/arrow-adbc/go/adbc/sqldriver"
)

const DriverName = "adbcduck"

func Make(libraryName string) *Driver {
	return &Driver{
		adbcDriver: sqldriver.Driver{
			Driver: drivermgr.Driver{},
		},
		libraryName: libraryName,
	}
}

const namePattern = "driver=%s;entrypoint=duckdb_adbc_init;path=%s"

type Driver struct {
	adbcDriver  sqldriver.Driver
	libraryName string
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	return d.adbcDriver.OpenConnector(d.getDsn(name))
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return d.adbcDriver.Open(d.getDsn(name))
}

func (d *Driver) getDsn(name string) string {
	return fmt.Sprintf(namePattern, d.libraryName, name)
}
