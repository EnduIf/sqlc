-- name: GetSensor :one
SELECT id, name, location, created_at FROM sensors
WHERE id = $1;

-- name: ListSensors :many
SELECT id, name, location, created_at FROM sensors
ORDER BY name;

-- name: CreateSensor :one
INSERT INTO sensors (
  id, name, location, created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

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