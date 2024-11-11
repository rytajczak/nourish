-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, email, provider, picture)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateSpoonCredential :one
INSERT INTO spoon_credential (user_id, username, password, hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetDailyGoal :one
SELECT calories, carbs, protein, fat FROM daily_goal WHERE user_id = $1;

-- name: UpdateDailyGoal :one
INSERT INTO daily_goal (user_id, calories, carbs, protein, fat)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id) 
DO UPDATE SET 
    calories = EXCLUDED.calories,
    carbs = EXCLUDED.carbs,
    protein = EXCLUDED.protein,
    fat = EXCLUDED.fat
RETURNING calories, carbs, protein, fat;

-- name: GetIntolerances :many
SELECT i.name
FROM users_intolerance ui
JOIN intolerance i ON i.id = ui.intolerance_id
WHERE ui.user_id = $1;

-- name: AddIntolerance :exec
INSERT INTO users_intolerance (user_id, intolerance_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveIntolerance :exec
DELETE FROM users_intolerance
WHERE user_id = $1 AND intolerance_id = $2;

-- name: GetDislikedIngredients :many
SELECT di.name
FROM users_disliked_ingredient udi
JOIN disliked_ingredient di ON di.id = udi.ingredient_id
WHERE udi.user_id = $1;

-- name: AddDislikedIngredient :one
WITH inserted_ingredient AS (
    INSERT INTO disliked_ingredient (id, name)
    VALUES (DEFAULT, $1)
    ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name
    RETURNING id
)
INSERT INTO users_disliked_ingredient (user_id, ingredient_id)
SELECT $2, id FROM inserted_ingredient
RETURNING *;

-- name: GetLikedRecipes :many
SELECT lr.*
FROM users_liked_recipe ulr
JOIN liked_recipe lr ON lr.id = ulr.recipe_id
WHERE ulr.user_id = $1;

-- name: AddLikedRecipe :one
WITH inserted_recipe AS (
    INSERT INTO liked_recipe (spoon_id, title, image, calories, protein, carbs, fat)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id
)
INSERT INTO users_liked_recipe (user_id, recipe_id)
SELECT $8, id FROM inserted_recipe
RETURNING *;
