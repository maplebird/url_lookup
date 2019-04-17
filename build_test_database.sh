#!/usr/bin/env bash

set -u
set -x

# Set MySQL root pw
MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD:-password}

# Run MySQL in docker
if ! docker images | grep mysql; then
    docker pull mysql
fi

# Delete existing docker container database
if docker ps -af name=url_lookup_db | grep url_lookup_db; then
    CONTAINER_ID=$(docker ps -a -f name=url_lookup_db | grep url_lookup_db | awk '{print $1}')
    docker stop ${CONTAINER_ID} || echo "Container already stopped"
    docker rm ${CONTAINER_ID}
fi

# Start new database instance
docker run --name url_lookup_db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -d mysql

# Give MySQL time to start
sleep 30

# Execute migrations
# Make sure `mysql` binary is installed on your system
# Should really be handled by flyway
echo "Testing database connection"
mysql -h 127.0.0.1 -uroot -ppassword -e "show databases;"

MIGRATIONS=$(ls src/migrations/*.sql)
for MIGRATION in ${MIGRATIONS}; do
    echo "Running database schema migration ${MIGRATION}"
    mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} < ${MIGRATION}
done