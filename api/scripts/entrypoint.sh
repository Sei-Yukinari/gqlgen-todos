#!/bin/bash

set -e

# 起動
arelo -p '**/*.go' -p '**/*.toml' -- go run .
