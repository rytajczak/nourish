-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (id, display_name, email, provider)
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING *;
