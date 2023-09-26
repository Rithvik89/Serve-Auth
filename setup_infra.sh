#!/bin/bash

# Define variables
CONTAINER_NAME="primary_db"
POSTGRES_USER="root"
POSTGRES_PASSWORD="mypassword"
DATABASE_NAME="authdb"
HOST="localhost"

# Delete containers if there are any with the existing name
make remove_db
# Run PostgreSQL container with specified user, password, and database
make create_db
# Wait for the container to start (adjust sleep time as needed)
#TODO: Dynamic adjustment of this time interval is to be done
sleep 10

# Connect to the container and create the database (you might need to install `psql` in your host system)
docker exec -it "$CONTAINER_NAME" psql -U "$POSTGRES_USER" -h "$HOST" -c "CREATE DATABASE $DATABASE_NAME;"

echo "PostgreSQL container '$CONTAINER_NAME' with database '$DATABASE_NAME' created successfully!"
