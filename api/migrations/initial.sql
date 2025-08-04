CREATE TABLE IF NOT EXISTS users
(
    id            bigserial PRIMARY KEY,
    name          varchar(128) NOT NULL,
    email         varchar(128) NOT NULL,
    password_hash bytea        NOT NULL,
    date_created  timestamp,
    date_modified timestamp
);

CREATE TABLE IF NOT EXISTS food_items
(
    id            bigserial PRIMARY KEY,
    manufacturer  varchar(128),
    product       varchar(128),
    protein       double precision,
    carbohydrate  double precision,
    fat           double precision,
    kcal          double precision,
    public        boolean,
    source        varchar(128),
    date_created  timestamp,
    date_modified timestamp,
    owner_id      bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS micronutrients
(
    id            bigserial PRIMARY KEY,
    name          varchar(128),
    amount        double precision,
    date_created  timestamp,
    date_modified timestamp,
    food_item_id  bigint REFERENCES food_items (id) ON DELETE CASCADE NOT NULL
);



CREATE TABLE IF NOT EXISTS meals
(
    id              bigserial PRIMARY KEY,
    sequence_number integer,
    meal_time       timestamp,
    date_created    timestamp,
    date_modified   timestamp,
    owner_id        bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS meal_entries
(
    id            bigserial PRIMARY KEY,
    amount        double precision                               NOT NULL,
    food_item_id  bigint REFERENCES food_items (id),
    recipe_id     bigint REFERENCES food_items (id),
    date_created  timestamp,
    date_modified timestamp,
    meal_id       bigint REFERENCES meals (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes
(
    id            bigserial PRIMARY KEY,
    name          varchar(128),
    date_created  timestamp,
    date_modified timestamp,
    owner_id      bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe_entries
(
    id            bigserial PRIMARY KEY,
    amount        double precision,
    food_item_id  bigint REFERENCES food_items (id),
    date_created  timestamp,
    date_modified timestamp,
    recipe_id     bigint REFERENCES recipes (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS portion_sizes
(
    id           bigserial PRIMARY KEY,
    name         varchar(128),
    amount       double precision,
    date_created  timestamp,
    date_modified timestamp,
    food_item_id bigint REFERENCES food_items (id),
    recipe_id    bigint REFERENCES recipes (id)
);

/* functions */
CREATE FUNCTION set_date_modified() RETURNS trigger AS
$$
BEGIN
    new.date_modified = NOW();
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION set_date_created() RETURNS trigger AS
$$
BEGIN
    new.date_created = NOW();
    new.date_modified = new.date_created;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

/* date_created when inserting */
CREATE TRIGGER set_created
    BEFORE INSERT
    ON recipe_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON recipes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON micronutrients
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON meal_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON food_items
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON meals
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_created
    BEFORE INSERT
    ON portion_sizes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();


/* date_modified when updating*/
CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON recipe_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON recipes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON micronutrients
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON meal_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON food_items
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON meals
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON portion_sizes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();
