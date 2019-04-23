#!/usr/bin/env bash

set -x

HERE=$(pwd)
export GOPATH=${HERE}

# Build
cd src/url_lookup
go get -v

# Run all tests, including integration tests
# Docker build does not run integration.
test() {
    cp integration_test.go_disable integration_test.go
    go test -v
    rm integration_test.go
}

echo "Running unit and integration tests"
if test; then
    echo "Tests succeeded.  Compiling the app."
    go build -o url_lookup
else
    echo "Tests failed.  Make sure to start database container first."
    echo "./build_and_start_db_container.sh"
fi

echo "Starting local server"
./url_lookup