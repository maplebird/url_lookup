#!/usr/bin/env bash

set -ue

# Build url_lookup container
echo "Building url_lookup container image"
docker build -t url_lookup_server .

