-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (username, email, provider, picture, diet, calories, carbs, protein, fat)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: CreateSpoonCredential :one
INSERT INTO spoon_credential (user_id, username, password, hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUserPreferences :one
UPDATE users SET diet = $1, calories = $2, protein = $3, carbs = $4, fat = $5 WHERE email = $6
RETURNING *;

-- name: GetUserIntolerances :many
SELECT i.name
FROM users_intolerance ui
JOIN intolerance i ON ui.intolerance_id = i.id
JOIN users u ON ui.user_id = u.id
WHERE u.email = $1;

-- name: CreateUserIntolerance :one
INSERT INTO users_intolerance (user_id, intolerance_id)
SELECT $1, i.id
FROM intolerance i
WHERE i.name = $2
RETURNING *;

-- name: DeleteUserIntolerance :one
DELETE FROM users_intolerance WHERE user_id = $1
RETURNING *;
