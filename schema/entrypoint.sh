

echo "Postgres is up - executing migrations"

psql -h "db" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /schema/000001_init.up.sql

exec postgres