#!/usr/bin/env bash

set -e

## Perform all actions that are run on TravisCI to test pipeline

# Installing Test
go get -t -v ./
go get -u github.com/gobuffalo/packr/v2/packr2

# Building Test
packr2
go build
go test ./test/

# Deployment Test

go mod tidy
packr2
go generate ./