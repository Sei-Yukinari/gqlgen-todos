#!/bin/bash

set -e

# db migration
goose up
echo "migrated."

# 起動
arelo -p '**/*.go' -p '**/*.toml' -- go run .
