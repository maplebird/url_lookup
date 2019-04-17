#!/usr/bin/env bash

# Delete MySQL container
if docker ps -af name=url_lookup_db | grep url_lookup_db; then
    CONTAINER_ID=$(docker ps -a -f name=url_lookup_db | grep url_lookup_db | awk '{print $1}')
    docker stop ${CONTAINER_ID} || echo "Container already stopped"
    docker rm ${CONTAINER_ID}
fi

rm -f src/main/main