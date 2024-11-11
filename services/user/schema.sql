CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    provider VARCHAR(50) NOT NULL,
    picture VARCHAR(100),
    diet VARCHAR(50) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE spoon_credential (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    hash VARCHAR(100) NOT NULL
);

CREATE TABLE daily_goal (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    calories INTEGER,
    carbs INTEGER,
    protein INTEGER,
    fat INTEGER
);

CREATE TABLE intolerance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE disliked_ingredient (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE liked_recipe (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    spoon_id INTEGER,
    title VARCHAR(100) NOT NULL,
    image VARCHAR(100),
    calories INTEGER,
    protein INTEGER,
    carbs INTEGER,
    fat INTEGER
);

CREATE TABLE users_intolerance (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    intolerance_id UUID REFERENCES intolerance(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, intolerance_id)
);

CREATE TABLE users_disliked_ingredient (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    ingredient_id INTEGER REFERENCES disliked_ingredient(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, ingredient_id)
);

CREATE TABLE users_liked_recipe (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    recipe_id UUID REFERENCES liked_recipe(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, recipe_id)
);

INSERT INTO intolerance (name) VALUES 
    ('dairy'),
    ('egg'),
    ('gluten'), 
    ('peanut'),
    ('sesame'),
    ('seafood'),
    ('shellfish'),
    ('soy'),
    ('sulfite'),
    ('tree nut'),
    ('wheat');