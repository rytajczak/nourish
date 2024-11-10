// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type DailyGoal struct {
	UserID   pgtype.UUID
	Calories pgtype.Int4
	Carbs    pgtype.Int4
	Protein  pgtype.Int4
	Fat      pgtype.Int4
}

type DislikedIngredient struct {
	ID   int32
	Name string
}

type Intolerance struct {
	ID   pgtype.UUID
	Name string
}

type LikedRecipe struct {
	ID       pgtype.UUID
	SpoonID  pgtype.Int4
	Title    string
	Image    pgtype.Text
	Calories pgtype.Int4
	Protein  pgtype.Int4
	Carbs    pgtype.Int4
	Fat      pgtype.Int4
}

type SpoonCredential struct {
	UserID   pgtype.UUID
	Username string
	Password string
	Hash     string
}

type User struct {
	ID         pgtype.UUID
	Username   string
	Email      string
	Provider   string
	Picture    pgtype.Text
	Diet       pgtype.Text
	CreatedAt  pgtype.Timestamp
	ModifiedAt pgtype.Timestamp
}

type UsersDislikedIngredient struct {
	UserID       pgtype.UUID
	IngredientID int32
}

type UsersIntolerance struct {
	UserID        pgtype.UUID
	IntoleranceID pgtype.UUID
}

type UsersLikedRecipe struct {
	UserID   pgtype.UUID
	RecipeID pgtype.UUID
}
