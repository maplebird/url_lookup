#!/usr/bin/env bash

# Build database and app containers
./build_and_start_db_container.sh
./build_app_container.sh
./start_app_container.sh
