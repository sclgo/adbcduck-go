# adbcduck-go - alternative DuckDB Go driver

`adbcduck-go` is a Go [database/sql](https://pkg.go.dev/database/sql) driver for [DuckDB](https://duckdb.org/)
[ADBC API](https://duckdb.org/docs/clients/adbc). It is an alternative to the [official Go driver](https://duckdb.org/docs/clients/go) `github.com/marcboeker/go-duckdb`.
`adbcduck-go` is a thin, but not trivial, wrapper over generic `database/sql` [adapter](https://pkg.go.dev/github.com/apache/arrow-adbc/go/adbc/sqldriver) 
for [ADBC](https://arrow.apache.org/adbc/) drivers, maintained by the [Apache Arrow project](https://arrow.apache.org/).

[![Go Reference](https://pkg.go.dev/badge/github.com/sclgo/adbcduck-go.svg)](https://pkg.go.dev/github.com/sclgo/adbcduck-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/sclgo/adbcduck-go)](https://goreportcard.com/report/github.com/sclgo/adbcduck-go)
[![Tests](https://github.com/sclgo/adbcduck-go/actions/workflows/go.yml/badge.svg)](https://coveralls.io/github/sclgo/adbcduck-go)

`adbcduck-go` is great for:

- **libraries that value their dependency footprint** - `go get github.com/sclgo/adbcduck-go` downloads up to 310 MB in Go modules,
  while `go get github.com/marcboeker/go-duckdb` downloads up to 824 MB - 500 MB more. Even when using GOPROXY, this quickly adds up
  especially in CI builds. Remember, projects that import your Go library must download all your transitive dependencies, even
  you use them only in tests.
- **developers that hit [issues](https://github.com/marcboeker/go-duckdb/issues) in the official library** - while the approach
  is this library is not inherently better (or worse) than the approach in the official Go client, it is unlikely that
  the two different codebases will have the same bugs. The Arrow ADBC Go library which implements the majority of this
  driver is [actively developed and supported](https://github.com/apache/arrow-adbc/pulse/monthly).
- **apps that need to work with a specific DuckDB version or even multiple version at the same time** - 
  This driver loads the DuckDB dynamic library at runtime. Multiple DuckDB dynamic libraries can be used 
  at the same time in the same app. In contrast, the official DuckDB Go client either works 
  with the fixed bundled DuckDB release or with a single specific dynamic 
  library, [specified](https://github.com/marcboeker/go-duckdb?tab=readme-ov-file#dynamic-linking) both at compile time and at runtime.
  - DuckDB guarantees backward compatibility of database files but 
    [not forward compatibility](https://duckdb.org/docs/stable/internals/storage.html#compatibility) so developers
    should carefully choose client library and DuckDB versions when using shared DBs.

## Quickstart

To get started quickly, use:

* blank import path - `github.com/sclgo/adbcduck-go/quickstart`
* driver name - `adbcduck`
* data source name - path to DB as specified in <https://duckdb.org/docs/stable/connect/overview>, 
  optionally followed by `key=value` [options](https://duckdb.org/docs/stable/configuration/overview#configuration-reference)
  separated by `;`
  * example: `test.db;threads=4`
  * use either empty path or `:memory:` to specify an in-memory database

```go
import (
	"database/sql"
	_ "github.com/sclgo/adbcduck-go/quickstart"
)

...

sql.Open("adbcduck", "") // Opens an in-memory database
```

## Configuration

Review [the API documentation](https://pkg.go.dev/github.com/sclgo/adbcduck-go)
and the minimal code of the [quickstart](/quickstart/quickstart.go) package to see how to configure the library name, 
location, and, optionally, automatic download.

Registration with blank import is optional. Importing only `github.com/sclgo/adbcduck-go` gives access to
the entire driver API without side effects. Only the `github.com/sclgo/adbcduck-go/register`
and `github.com/sclgo/adbcduck-go/quickstart` packages have `init()` functions that call 
[sql.Register](https://pkg.go.dev/database/sql#Register) .

## CLI

`adbcduck-go` is compatible with [usql](https://github.com/xo/usql) - the universal SQL CLI - a fully featured, single-binary, 
CLI for over 50 databases. Check out [this example](./docs/cli_example.md).

## Limitations

This library requires CGO, because it depends on the [CGO wrapper](https://github.com/apache/arrow-adbc/blob/11a9128/go/adbc/drivermgr/wrapper.go)
for the [ADBC Driver Manager](https://arrow.apache.org/adbc/main/cpp/driver_manager.html).
In the future, the CGO wrapper may be replaced with one using <https://github.com/ebitengine/purego>. 
