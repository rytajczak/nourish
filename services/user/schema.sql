CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    provider VARCHAR(50) NOT NULL,
    picture VARCHAR(100),
    diet VARCHAR(50) DEFAULT 'none',
    calories INTEGER DEFAULT 0,
    carbs INTEGER DEFAULT 0,
    protein INTEGER DEFAULT 0,
    fat INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE spoon_credential (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    hash VARCHAR(100) NOT NULL
);

CREATE TABLE intolerance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE liked_recipe (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    spoon_id INTEGER,
);

CREATE TABLE users_intolerance (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    intolerance_id UUID REFERENCES intolerance(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, intolerance_id)
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