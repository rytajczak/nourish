// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSpoonCredential = `-- name: CreateSpoonCredential :one
INSERT INTO spoon_credential (user_id, username, password, hash)
VALUES ($1, $2, $3, $4)
RETURNING user_id, username, password, hash
`

type CreateSpoonCredentialParams struct {
	UserID   pgtype.UUID `json:"user_id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Hash     string      `json:"hash"`
}

func (q *Queries) CreateSpoonCredential(ctx context.Context, arg CreateSpoonCredentialParams) (SpoonCredential, error) {
	row := q.db.QueryRow(ctx, createSpoonCredential,
		arg.UserID,
		arg.Username,
		arg.Password,
		arg.Hash,
	)
	var i SpoonCredential
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Hash,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, provider, picture, diet, calories, carbs, protein, fat)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, username, email, provider, picture, diet, calories, carbs, protein, fat, created_at, modified_at
`

type CreateUserParams struct {
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Provider string      `json:"provider"`
	Picture  pgtype.Text `json:"picture"`
	Diet     pgtype.Text `json:"diet"`
	Calories pgtype.Int4 `json:"calories"`
	Carbs    pgtype.Int4 `json:"carbs"`
	Protein  pgtype.Int4 `json:"protein"`
	Fat      pgtype.Int4 `json:"fat"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Provider,
		arg.Picture,
		arg.Diet,
		arg.Calories,
		arg.Carbs,
		arg.Protein,
		arg.Fat,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Provider,
		&i.Picture,
		&i.Diet,
		&i.Calories,
		&i.Carbs,
		&i.Protein,
		&i.Fat,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const createUserIntolerance = `-- name: CreateUserIntolerance :one
INSERT INTO users_intolerance (user_id, intolerance_id)
SELECT $1, i.id
FROM intolerance i
WHERE i.name = $2
RETURNING user_id, intolerance_id
`

type CreateUserIntoleranceParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Name   string      `json:"name"`
}

func (q *Queries) CreateUserIntolerance(ctx context.Context, arg CreateUserIntoleranceParams) (UsersIntolerance, error) {
	row := q.db.QueryRow(ctx, createUserIntolerance, arg.UserID, arg.Name)
	var i UsersIntolerance
	err := row.Scan(&i.UserID, &i.IntoleranceID)
	return i, err
}

const deleteUserIntolerance = `-- name: DeleteUserIntolerance :one
DELETE FROM users_intolerance WHERE user_id = $1
RETURNING user_id, intolerance_id
`

func (q *Queries) DeleteUserIntolerance(ctx context.Context, userID pgtype.UUID) (UsersIntolerance, error) {
	row := q.db.QueryRow(ctx, deleteUserIntolerance, userID)
	var i UsersIntolerance
	err := row.Scan(&i.UserID, &i.IntoleranceID)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, provider, picture, diet, calories, carbs, protein, fat, created_at, modified_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Provider,
		&i.Picture,
		&i.Diet,
		&i.Calories,
		&i.Carbs,
		&i.Protein,
		&i.Fat,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, email, provider, picture, diet, calories, carbs, protein, fat, created_at, modified_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Provider,
		&i.Picture,
		&i.Diet,
		&i.Calories,
		&i.Carbs,
		&i.Protein,
		&i.Fat,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUserIntolerances = `-- name: GetUserIntolerances :many
SELECT i.name
FROM users_intolerance ui
JOIN intolerance i ON ui.intolerance_id = i.id
JOIN users u ON ui.user_id = u.id
WHERE u.email = $1
`

func (q *Queries) GetUserIntolerances(ctx context.Context, email string) ([]string, error) {
	rows, err := q.db.Query(ctx, getUserIntolerances, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserProfile = `-- name: GetUserProfile :one
SELECT diet, calories, protein, carbs, fat FROM users WHERE email = $1
`

type GetUserProfileRow struct {
	Diet     pgtype.Text `json:"diet"`
	Calories pgtype.Int4 `json:"calories"`
	Protein  pgtype.Int4 `json:"protein"`
	Carbs    pgtype.Int4 `json:"carbs"`
	Fat      pgtype.Int4 `json:"fat"`
}

func (q *Queries) GetUserProfile(ctx context.Context, email string) (GetUserProfileRow, error) {
	row := q.db.QueryRow(ctx, getUserProfile, email)
	var i GetUserProfileRow
	err := row.Scan(
		&i.Diet,
		&i.Calories,
		&i.Protein,
		&i.Carbs,
		&i.Fat,
	)
	return i, err
}

const getUsernameAndHash = `-- name: GetUsernameAndHash :one
SELECT username, hash FROM spoon_credential WHERE user_id = $1
`

type GetUsernameAndHashRow struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

func (q *Queries) GetUsernameAndHash(ctx context.Context, userID pgtype.UUID) (GetUsernameAndHashRow, error) {
	row := q.db.QueryRow(ctx, getUsernameAndHash, userID)
	var i GetUsernameAndHashRow
	err := row.Scan(&i.Username, &i.Hash)
	return i, err
}

const updateUserProfile = `-- name: UpdateUserProfile :one
UPDATE users SET diet = $1, calories = $2, protein = $3, carbs = $4, fat = $5 WHERE email = $6
RETURNING diet, calories, protein, carbs, fat
`

type UpdateUserProfileParams struct {
	Diet     pgtype.Text `json:"diet"`
	Calories pgtype.Int4 `json:"calories"`
	Protein  pgtype.Int4 `json:"protein"`
	Carbs    pgtype.Int4 `json:"carbs"`
	Fat      pgtype.Int4 `json:"fat"`
	Email    string      `json:"email"`
}

type UpdateUserProfileRow struct {
	Diet     pgtype.Text `json:"diet"`
	Calories pgtype.Int4 `json:"calories"`
	Protein  pgtype.Int4 `json:"protein"`
	Carbs    pgtype.Int4 `json:"carbs"`
	Fat      pgtype.Int4 `json:"fat"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (UpdateUserProfileRow, error) {
	row := q.db.QueryRow(ctx, updateUserProfile,
		arg.Diet,
		arg.Calories,
		arg.Protein,
		arg.Carbs,
		arg.Fat,
		arg.Email,
	)
	var i UpdateUserProfileRow
	err := row.Scan(
		&i.Diet,
		&i.Calories,
		&i.Protein,
		&i.Carbs,
		&i.Fat,
	)
	return i, err
}
