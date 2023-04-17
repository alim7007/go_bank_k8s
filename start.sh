#!/bin/sh

set -e

echo "run db migration"
# source /app/.env
/app/migrate -path /app/migration -database "postgresql://root:olim123@localhost:5432/olim_bank?sslmode=disable" -verbose up

echo "start the app"
exec "$@"