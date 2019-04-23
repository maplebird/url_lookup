#!/usr/bin/env bash

set -eu

HERE=$(pwd)
export GOPATH=${HERE}

# Clean up old builds
rm -rf bin

# Build
cd src/url_lookup
go get -v

# Run all tests, including integration tests
# Docker build does not run integration tests
test() {
    cp testPackage02_integration_test.go_disable testPackage02_integration_test.go
    go test -v .
    rm -f testPackage02_integration_test.go
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