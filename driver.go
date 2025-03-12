package adbcduck // import "github.com/sclgo/duckdb-adbc-go"

import (
	"context"
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

type duckDbConnector struct {
	driver.Connector
}

func (d duckDbConnector) Connect(ctx context.Context) (driver.Conn, error) {
	adbcConn, err := d.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &conn{
		fullConn: adbcConn.(fullConn),
	}, nil
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	adbcCtr, err := d.adbcDriver.OpenConnector(d.getDsn(name))
	return duckDbConnector{adbcCtr}, err
}

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
