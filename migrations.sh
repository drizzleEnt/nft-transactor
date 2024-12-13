#!/bin/bash
source .env

export MIGRATION_DSN="host=pg_nft port=5432 dbname=$DB_NAME user=$DB_USER password=$DB_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v