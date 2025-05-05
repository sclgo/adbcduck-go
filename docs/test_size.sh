#!/bin/bash

set -e

du -sh $(go env GOMODCACHE) || true
mkdir test && cd test
go mod init test
echo "package test \
\
import _ \"$PKG\" \
" >> import.go
go mod tidy
du -sh $(go env GOMODCACHE)
