CREATE TABLE recipe (
    id UUID PRIMARY KEY,
    spoonacular_id INTEGER UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    instructions TEXT,
    prep_time INTEGER,
    cook_time INTEGER,
    servings INTEGER,
    calories INTEGER,
    protein FLOAT,
    carbs FLOAT,
    fat FLOAT,
    image_url VARCHAR(255),
    source_url VARCHAR(255),
    is_custom BOOLEAN DEFAULT false,
    created_by UUID,  -- No foreign key, just store the user ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ingredient (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    unit VARCHAR(20)
);

CREATE TABLE recipe_ingredient (
    recipe_id UUID REFERENCES recipe(id) ON DELETE CASCADE,
    ingredient_id UUID REFERENCES ingredient(id) ON DELETE CASCADE,
    amount FLOAT NOT NULL,
    unit VARCHAR(20),
    PRIMARY KEY (recipe_id, ingredient_id)
);