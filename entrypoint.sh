#!/bin/bash

set -e

# RDB接続確認
until mysqladmin ping -h "${DB_HOST}" -P "${DB_PORT}" --silent; do
  echo "waiting for mysql..."
  sleep 2
done
echo "success to connect mysql."

# testRDB作成
if [ "$GO_ENV" = "development" ]; then
  mysql -h "${DB_HOST}" -P "${DB_PORT}" -u "${DB_USER}" -p"${DB_PASSWORD}" -e "CREATE DATABASE IF NOT EXISTS "$TEST_DB_NAME";"
fi

if [ "$GO_ENV" = "development" ]; then
  # arelo -p '**/*.go' -p '**/*.toml' -- go run ./main.go
  arelo -p '**/*.go' -- go run ./main.go
fi
