#!/usr/bin/env bash

# Delete MySQL and server containers
echo "Stopping and deleting MySQL and server containers"
if docker ps -af name=url_lookup | grep url_lookup; then
    CONTAINER_IDS=$(docker ps -a -f name=url_lookup | grep url_lookup | awk '{print $1}')

    for ID in ${CONTAINER_IDS}; do
        docker stop ${ID} || echo "Container already stopped"
        docker rm ${ID}
    done
fi

# Remove compiled binary
rm -f src/main/main