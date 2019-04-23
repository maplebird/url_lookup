#!/usr/bin/env bash

set -ue

# Set MySQL root pw
MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD:-password}
MYSQL_HOST_PORT=${URL_LOOKUP_DBPORT-:3306}

# Run MySQL in docker
if ! docker images | grep mysql; then
    echo "MySQL docker image not found; downloading"
    if ! docker pull mysql; then
        echo "Cannot pull mysql image; make sure to run docker login first"
        exit 1
    fi
fi

# Delete existing docker container database
echo "Deleting existing url_lookup_db database"
if docker ps -af name=url_lookup_db | grep url_lookup_db; then
    CONTAINER_ID=$(docker ps -a -f name=url_lookup_db | grep url_lookup_db | awk '{print $1}')
    docker stop ${CONTAINER_ID} || echo "Container already stopped"
    docker rm ${CONTAINER_ID}
fi

# Start new database instance
docker run --name url_lookup_db -p ${MYSQL_HOST_PORT}:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -d mysql

# Give MySQL time to start
sleep 30

# Execute migrations
# Make sure `mysql` binary is installed on your system
# Should really be handled by flyway
echo "Testing database connection."
if ! mysql -h 127.0.0.1 -P ${MYSQL_HOST_PORT} -u root -p${MYSQL_ROOT_PASSWORD} -e "show databases;"; then
    echo "Could not connect to database.  Most likely container did not start in time."
    echo "Try increasing timeout above to a larger value, like 60 seconds."
    exit 1
else
    echo "Successfully connected to test database.  Executing migrations."
    MIGRATIONS=$(ls migrations/*.sql)
    for MIGRATION in ${MIGRATIONS}; do
        echo "Running database schema migration ${MIGRATION}"
        mysql -h 127.0.0.1 -P ${MYSQL_HOST_PORT} -uroot -p${MYSQL_ROOT_PASSWORD} < ${MIGRATION}
    done
fi
echo "Database migrations completed successfully."
