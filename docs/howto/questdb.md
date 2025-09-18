# Using sqlc with QuestDB

[QuestDB](https://questdb.io/) is a fast open-source time-series database for high throughput ingestion and fast SQL queries with extensions for time-series analysis.

sqlc supports QuestDB using the `questdb` engine. QuestDB is largely PostgreSQL-compatible, so sqlc treats QuestDB queries similar to PostgreSQL with some specific type mappings.

## Configuration

To use QuestDB, set the `engine` field to `questdb` in your `sqlc.yaml` configuration:

```yaml
version: "2"
sql:
  - engine: "questdb"
    queries: "query.sql"
    schema: "schema.sql" 
    gen:
      go:
        package: "db"
        out: "."
```

## Schema

QuestDB supports standard SQL DDL with some extensions for time-series functionality. Here's an example schema:

```sql
-- Standard table creation
CREATE TABLE sensors (
    id uuid NOT NULL,
    name text NOT NULL,
    location text,
    created_at timestamp
);

-- Time-series table
CREATE TABLE measurements (
    sensor_id uuid NOT NULL,
    value double precision NOT NULL,
    temperature double precision,
    humidity int,
    ts timestamp,
    metadata text
);
```

Note: QuestDB-specific syntax like `timestamp(ts) PARTITION BY DAY` should be handled through migrations or database-specific DDL, as sqlc focuses on the portable SQL schema definition.

## Type Mappings

QuestDB uses PostgreSQL-compatible type mappings with a few extensions:

| QuestDB Type | Go Type |
|-------------|----------|
| `symbol` | `string` / `sql.NullString` |
| `geohash` | `string` / `sql.NullString` |
| All other types | Same as PostgreSQL |

## Queries

Write your queries using standard SQL syntax:

```sql
-- name: GetSensor :one
SELECT id, name, location, created_at FROM sensors
WHERE id = $1;

-- name: GetMeasurements :many
SELECT sensor_id, value, temperature, humidity, ts, metadata
FROM measurements
WHERE sensor_id = $1
ORDER BY ts DESC
LIMIT 100;

-- name: CreateMeasurement :exec
INSERT INTO measurements (
  sensor_id, value, temperature, humidity, ts, metadata
) VALUES (
  $1, $2, $3, $4, $5, $6
);
```

QuestDB-specific time-series functions and syntax can be used in your queries as they are passed through to the database.

## Example

See the [QuestDB example](https://github.com/EnduIf/sqlc/tree/main/examples/questdb) for a complete working example.