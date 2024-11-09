CREATE TABLE meal_plan (
    id UUID PRIMARY KEY,
    user_id UUID,
    week_start_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE meal_plan_recipe (
    meal_plan_id UUID REFERENCES meal_plan(id) ON DELETE CASCADE,
    recipe_id UUID,
    day_of_week INTEGER CHECK (day_of_week BETWEEN 0 AND 6),
    meal_type VARCHAR(20) CHECK (meal_type IN ('breakfast', 'lunch', 'dinner', 'snack')),
    servings INTEGER DEFAULT 1,
    PRIMARY KEY (meal_plan_id, recipe_id, day_of_week, meal_type)
);

CREATE TABLE grocery_list (
    id UUID PRIMARY KEY,
    user_id UUID,
    meal_plan_id UUID REFERENCES meal_plan(id) ON DELETE CASCADE,
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE grocery_list_item (
    list_id UUID REFERENCES grocery_list(id) ON DELETE CASCADE,
    ingredient_id UUID,
    amount FLOAT,
    unit VARCHAR(20),
    is_checked BOOLEAN DEFAULT false,
    PRIMARY KEY (list_id, ingredient_id)
);