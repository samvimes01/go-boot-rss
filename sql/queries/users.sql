-- name: CreateUser :one
INSERT INTO users (id, name, api_key)
VALUES ($1, $2, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;
