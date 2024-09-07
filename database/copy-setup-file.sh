#!/bin/sh

if [ "$1" = "development" ]; then
  cp /app/setup-database-dev.sql /docker-entrypoint-initdb.d/
else
  cp /app/setup-database-stable.sql /docker-entrypoint-initdb.d/
fi
