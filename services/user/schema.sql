CREATE TABLE users (
    id UUID PRIMARY KEY,
    display_name VARCHAR(50),
    email VARCHAR(255) NOT NULL UNIQUE,
    provider VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_sign_in_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profile (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    username VARCHAR(50) NOT NULL,
    picture VARCHAR(100),
    diet VARCHAR(100)
);

CREATE TABLE security (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    spoonacular_username VARCHAR(100),
    spoonacular_hash VARCHAR(100),
    spoonacular_password VARCHAR(100)
);

CREATE TABLE intolerance (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE saved_recipe (
    id int PRIMARY KEY
);

CREATE TABLE profile_intolerance (
    profile_id UUID REFERENCES profile(user_id) ON DELETE CASCADE,
    intolerance_id UUID REFERENCES intolerance(id) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, intolerance_id)
);

CREATE TABLE profile_saved_recipe (
    profile_id UUID REFERENCES profile(user_id) ON DELETE CASCADE,
    saved_recipe_id int REFERENCES saved_recipe(id) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, saved_recipe_id)
);

INSERT INTO intolerance (id, name) VALUES
  (gen_random_uuid(), 'dairy'),
  (gen_random_uuid(), 'egg'),
  (gen_random_uuid(), 'gluten'),
  (gen_random_uuid(), 'peanut'),
  (gen_random_uuid(), 'sesame'),
  (gen_random_uuid(), 'seafood'),
  (gen_random_uuid(), 'shellfish'),
  (gen_random_uuid(), 'soy'),
  (gen_random_uuid(), 'sulfite'),
  (gen_random_uuid(), 'tree nut'),
  (gen_random_uuid(), 'wheat');
