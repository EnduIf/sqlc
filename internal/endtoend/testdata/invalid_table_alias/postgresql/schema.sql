-- https://github.com/EnduIf/sqlc/issues/437
CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
