# adbcduck-go - alternative DuckDB Go driver

`duckdb-adbc-go` is a Go [database/sql](https://pkg.go.dev/database/sql) driver for [duckdb](https://duckdb.org/)
 [ADBC API](https://duckdb.org/docs/clients/adbc). It is an altenative to the [official Go driver](https://duckdb.org/docs/clients/go) 
- `github.com/marcboeker/go-duckdb`.
`duckdb-adbc-go` is a very thin wrapper over generic `database/sql` [adapter](https://pkg.go.dev/github.com/apache/arrow-adbc/go/adbc/sqldriver) 
for [ADBC](https://arrow.apache.org/adbc/) drivers, maintained by the [Apache Arrow project](https://arrow.apache.org/).

`duckdb-adbc-go` is great for:

- **libraries that value their dependency footprint** - `go get github.com/sclgo/duckdb-adbc-go` downloads up to 310 MB in Go modules,
  while `go get github.com/marcboeker/go-duckdb` downloads up to 824 MB - 500 MB more. Even when using GOPROXY, this quickly adds up
  especially in CI builds. Remember, projects that import your Go library must download all your transitive dependencies, even
  you use them only in tests.
- **developers that hit [issues](https://github.com/marcboeker/go-duckdb/issues) in the official library** - while the approach
  is this library is not inherently better (or worse) than the approach in the official Go client, it is unlikely that
  the two different codebases will have the same bugs.
- **apps that need to work with multiple DuckDB versions** - this driver loads the DuckDB dynamic library at runtime.
  Multiple DuckDB dynamic libraries can be used at the same time in the same app. In constrast, the 
  official DuckDB Go client either works with the fixed bundled DuckDB release or with a single specific dynamic 
  library, [specified](https://github.com/marcboeker/go-duckdb?tab=readme-ov-file#dynamic-linking) both at compile time and at runtime.


