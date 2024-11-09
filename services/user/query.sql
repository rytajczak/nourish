-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (id, display_name, email, provider, created_at, last_sign_in_at)
VALUES (uuid_generate_v4(), $1, $2, $3, now(), now())
RETURNING *;

-- name: CreateProfile :one
INSERT INTO profile (user_id, username, picture, diet)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateSecurity :one
INSERT INTO security (user_id, spoonacular_username, spoonacular_hash, spoonacular_password)
VALUES ($1, $2, $3, $4)
RETURNING *;