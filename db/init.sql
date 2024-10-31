CREATE TABLE auth (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    provider VARCHAR(50),
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_sign_in_at TIMESTAMP
);

CREATE TABLE profile (
    id UUID PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    picture VARCHAR(100),
    spoonacular_password VARCHAR(100),
    hash VARCHAR(100),
    diet VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES auth(id) ON DELETE CASCADE
);

CREATE TABLE intolerance (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE liked_recipe (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profile_intolerance (
    profile_id UUID REFERENCES profile(id) ON DELETE CASCADE,
    intolerance_id UUID REFERENCES intolerance(id) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, intolerance_id)
);

CREATE TABLE profile_liked_recipe (
    profile_id UUID REFERENCES profile(id) ON DELETE CASCADE,
    liked_recipe_id UUID REFERENCES liked_recipe(id) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, liked_recipe_id)
);
