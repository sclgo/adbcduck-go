// Package sqldriver implements an adapter between ADBC Go API and database/sql/driver
// with changes specific to DuckDB
package sqldriver

// Most files in this package are forked from https://github.com/apache/arrow-adbc/tree/854d31e26/go/adbc/sqldriver
// under the Apache 2 license. Forked files have the Apache 2 header at the top.
// The commit ref in the above link will be updated as changes from upstream are integrated.

// The fork has the following changes compared to upstream:
// - transaction control is replaced with BEGIN, COMMIT, ROLLBACK statements; DuckDB's
//   support for the upstream approach - based on connection properties - is incomplete.
//   The transaction isolation ADBC property is not supported by DuckDB as of 1.2.0.
// - Driver.Open and OpenConnector accept DuckDB DSNs, as opposed to ADBC Driver DSNs.
//   DuckDB DSNs is converted to ADBC Driver Manager DSNs by Driver.getDsn.
// - experimental support for reading list types is added. Maybe offered upstream when it becomes more matures.

// The code is otherwise unchanged to make it easier to keep the fork up to date.
// Notably, types and fields which were exported upstream are exported here as well,
// so we can avoid capitalization differences to upstream.
// The package is in internal/ so that we don't export symbols that kept exported only for consistency.
