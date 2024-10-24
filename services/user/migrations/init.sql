CREATE TABLE auth (
    id VARCHAR PRIMARY KEY,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    provider VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_sign_in_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profile (
    id VARCHAR PRIMARY KEY,
    username VARCHAR NOT NULL,
    first_name VARCHAR,
    last_name VARCHAR,
    diet VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_auth FOREIGN KEY (id) REFERENCES auth(id) ON DELETE CASCADE
);

CREATE TABLE intolerence (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE liked_recipe (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    image VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profile_intolerence (
    profile_id VARCHAR,
    intolerence_id VARCHAR,
    PRIMARY KEY (profile_id, intolerence_id),
    CONSTRAINT fk_profile FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    CONSTRAINT fk_intolerence FOREIGN KEY (intolerence_id) REFERENCES intolerence(id) ON DELETE CASCADE
);

CREATE TABLE profile_liked_recipe (
    profile_id VARCHAR,
    liked_recipe_id VARCHAR,
    PRIMARY KEY (profile_id, liked_recipe_id),
    CONSTRAINT fk_profile FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    CONSTRAINT fk_liked_recipe FOREIGN KEY (liked_recipe_id) REFERENCES liked_recipe(id) ON DELETE CASCADE
);

-- modifed trigger
CREATE OR REPLACE FUNCTION update_modified_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.modified_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_auth_modified_at
BEFORE UPDATE ON auth
FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

CREATE TRIGGER trigger_profile_modified_at
BEFORE UPDATE ON profile
FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

CREATE TRIGGER trigger_liked_recipe_modified_at
BEFORE UPDATE ON liked_recipe
FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();