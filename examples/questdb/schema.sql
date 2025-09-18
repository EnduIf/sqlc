-- QuestDB schema example with time-series data
CREATE TABLE sensors (
    id uuid NOT NULL,
    name text NOT NULL,
    location text,
    created_at timestamp
);

CREATE TABLE measurements (
    sensor_id uuid NOT NULL,
    value double precision NOT NULL,
    temperature double precision,
    humidity int,
    ts timestamp,
    metadata text
);