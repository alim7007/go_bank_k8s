#!/bin/sh

set -e

echo "run db migration!!!! $DB_SOURCE"
echo "Docker_Entry_Point"
# #source /app/.env # for k8s
# /app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"




