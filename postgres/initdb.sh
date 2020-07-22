#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 -U postgres <<-EOSQL
    CREATE USER baristaschool WITH PASSWORD '$POSTGRES_PASSWORD';
    CREATE DATABASE baristaschool;
    GRANT ALL PRIVILEGES ON DATABASE baristaschool TO baristaschool;
EOSQL
