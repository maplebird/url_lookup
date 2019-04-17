#!/usr/bin/env bash

set -x

HERE=$(pwd)

# Build
cd src/main
go build

./main >> ${HERE}/server.log 2>&1