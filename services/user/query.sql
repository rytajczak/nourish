-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, email, provider, picture)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUserDiet :one
UPDATE users 
SET diet = $2, modified_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: CreateSpoonCredential :one
INSERT INTO spoon_credential (user_id, username, password, hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateDailyGoal :one
INSERT INTO daily_goal (user_id, calories, carbs, protein, fat)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id) 
DO UPDATE SET 
    calories = EXCLUDED.calories,
    carbs = EXCLUDED.carbs,
    protein = EXCLUDED.protein,
    fat = EXCLUDED.fat
RETURNING *;

-- name: GetUserIntolerances :many
SELECT i.name
FROM users_intolerance ui
JOIN intolerance i ON i.id = ui.intolerance_id
WHERE ui.user_id = $1;

-- name: AddUserIntolerance :exec
INSERT INTO users_intolerance (user_id, intolerance_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveUserIntolerance :exec
DELETE FROM users_intolerance
WHERE user_id = $1 AND intolerance_id = $2;

-- name: GetUserDislikedIngredients :many
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

-- name: GetUserLikedRecipes :many
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

-- name: GetUserWithPreferences :one
SELECT 
    u.*,
    sc.username as spoon_username,
    dg.calories,
    dg.carbs,
    dg.protein,
    dg.fat,
    ARRAY_AGG(DISTINCT i.name) as intolerances
FROM users u
LEFT JOIN spoon_credential sc ON u.id = sc.user_id
LEFT JOIN daily_goal dg ON u.id = dg.user_id
LEFT JOIN users_intolerance ui ON u.id = ui.user_id
LEFT JOIN intolerance i ON ui.intolerance_id = i.id
WHERE u.id = $1
GROUP BY u.id, sc.username, dg.calories, dg.carbs, dg.protein, dg.fat;
