#!/bin/bash

set -e

# db migration
migrate -source file://db/migrations -database \
 "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_DATABASE" \
  up
echo "migrated."

# 起動
arelo -p '**/*.go' -p '**/*.toml' -- go run .
