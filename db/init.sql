CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    provider VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_sign_in_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE security (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    spoonacular_hash VARCHAR(100),
    spoonacular_password VARCHAR(100)
);

CREATE TABLE profile (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    picture VARCHAR(100),
    diet VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE intolerance (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE liked_recipe (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profile_intolerance (
    profile_id UUID REFERENCES profile(user_id) ON DELETE CASCADE,
    intolerance_id UUID REFERENCES intolerance(id) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, intolerance_id)
);

CREATE TABLE profile_liked_recipe (
    profile_id UUID REFERENCES profile(user_id) ON DELETE CASCADE,
    liked_recipe_id UUID REFERENCES liked_recipe(id) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, liked_recipe_id)
);
