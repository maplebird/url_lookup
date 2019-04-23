#!/usr/bin/env bash

set -ue

# Build database and app containers
./build_and_start_db_container.sh
./build_app_container.sh

# If database is running on non-standard port, pass URL_LOOKUP_DBPORT env variable to container
URL_LOOKUP_DBPORT=${URL_LOOKUP_DBPORT:-3306}
DBPORT_DOCKER_ENV_VAR_STRING=""

if [[ ${URL_LOOKUP_DBPORT} -ne 3306 ]]; then
    DBPORT_DOCKER_ENV_VAR_STRING="--env URL_LOOKUP_DBPORT=${URL_LOOKUP_DBPORT}"
fi

# Get local IPv4 that's not localhost
# So server container can talk to a separate database container
LOCAL_ADDRESS=$(ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | awk '{print $2}' | grep -v '127.' | head -n1)

# Stop previous instance of this container
echo "Stopping url_lookup_server container if running"
if docker ps -af name=url_lookup_server | grep url_lookup_server; then
    CONTAINER_ID=$(docker ps -a -f name=url_lookup_server | grep url_lookup_server | awk '{print $1}')
    docker stop ${CONTAINER_ID} || echo "Container already stopped"
    docker rm ${CONTAINER_ID}
fi

echo "Starting new url_lookup_server container"
docker run \
    --name url_lookup_server \
    -p 5000:5000 \
    --env URL_LOOKUP_DBHOST=${LOCAL_ADDRESS} \
    ${DBPORT_DOCKER_ENV_VAR_STRING}  \
    -it url_lookup_server \
    /go/bin/url_lookup