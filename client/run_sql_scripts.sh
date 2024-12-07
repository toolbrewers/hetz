#!/bin/bash

# Database file path
DB_PATH="./db/client.db"

# Check if database file exists
if [ ! -f "$DB_PATH" ]; then
    echo "Creating database file..."
    touch "$DB_PATH"
fi

# Run all .sql files from the mock directory
for sql_file in ./db/mock/*.sql; do
    if [ -f "$sql_file" ]; then
        echo "Executing $sql_file..."
        sqlite3 "$DB_PATH" < "$sql_file"
    fi
done

# Clear the WAL file
echo "Clearing WAL file..."
sqlite3 "$DB_PATH" "PRAGMA wal_checkpoint;"

echo "All SQL scripts executed successfully!" 