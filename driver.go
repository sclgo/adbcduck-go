package duckadbc // import "github.com/sclgo/github-adbc-go"

import (
	"database/sql"
	"database/sql/driver"
	"os"

	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	"github.com/apache/arrow-adbc/go/adbc/sqldriver"
)

func init() {
	driverName := os.Getenv("DUCKDB_DRIVER_NAME")
	if driverName == "" {
		driverName = "duckadbc"
	}
	sql.Register(driverName, makeDriver())
}

func makeDriver() driver.Driver {
	return duckdbDriver{
		adbcDriver: sqldriver.Driver{
			Driver: drivermgr.Driver{},
		},
	}
}

const namePrefix = "driver=duckdb;entrypoint=duckdb_adbc_init;path="

type duckdbDriver struct {
	adbcDriver sqldriver.Driver
}

func (d duckdbDriver) OpenConnector(name string) (driver.Connector, error) {
	return d.adbcDriver.OpenConnector(namePrefix + name)
}

func (d duckdbDriver) Open(name string) (driver.Conn, error) {
	return d.adbcDriver.Open(namePrefix + name)
}
