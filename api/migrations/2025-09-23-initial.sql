CREATE TABLE users
(
    id            bigserial PRIMARY KEY,
    name          varchar(128) NOT NULL,
    email         varchar(128) NOT NULL,
    password_hash bytea        NOT NULL,
    date_created  timestamp,
    date_modified timestamp
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE food_items
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

CREATE TRIGGER set_created
    BEFORE INSERT
    ON food_items
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON food_items
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE food_item_micronutrients
(
    id            bigserial PRIMARY KEY,
    name          varchar(128),
    amount        double precision,
    date_created  timestamp,
    date_modified timestamp,
    food_item_id  bigint REFERENCES food_items (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON food_item_micronutrients
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON food_item_micronutrients
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE food_item_portion_sizes
(
    id            bigserial PRIMARY KEY,
    name          varchar(128),
    amount        double precision NOT NULL,
    date_created  timestamp,
    date_modified timestamp,
    food_item_id  bigint REFERENCES food_items (id) ON DELETE CASCADE
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON food_item_portion_sizes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON food_item_portion_sizes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE recipes
(
    id            bigserial PRIMARY KEY,
    name          varchar(128),
    date_created  timestamp,
    date_modified timestamp,
    owner_id      bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON recipes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON recipes
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


--The thinking here is that other entry types can be added in the future
CREATE TABLE recipe_food_item_entries
(
    id            bigserial PRIMARY KEY,
    amount        double precision,
    food_item_id  bigint REFERENCES food_items (id),
    date_created  timestamp,
    date_modified timestamp,
    recipe_id     bigint REFERENCES recipes (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON recipe_food_item_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON recipe_food_item_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE meals
(
    id              bigserial PRIMARY KEY,
    sequence_number integer,
    meal_time       timestamp,
    date_created    timestamp,
    date_modified   timestamp,
    owner_id        bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON meals
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON meals
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE meal_macronutrient_entries
(
    id              bigserial PRIMARY KEY,
    protein         double precision,
    carbohydrate    double precision,
    fat             double precision,
    kcal            double precision,
    sequence_number integer                                        NOT NULL,
    date_created    timestamp,
    date_modified   timestamp,
    meal_id         bigint REFERENCES meals (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON meal_macronutrient_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON meal_macronutrient_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE meal_recipe_entries
(
    id              bigserial PRIMARY KEY,
    recipe_id       bigint REFERENCES recipes (id)                 NOT NULL,
    amount          double precision                               NOT NULL,
    sequence_number integer                                        NOT NULL,
    date_created    timestamp,
    date_modified   timestamp,
    meal_id         bigint REFERENCES meals (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON meal_recipe_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON meal_recipe_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();


CREATE TABLE meal_food_item_entries
(
    id              bigserial PRIMARY KEY,
    food_item_id    bigint REFERENCES food_items (id)              NOT NULL,
    amount          double precision                               NOT NULL,
    sequence_number integer                                        NOT NULL,
    date_created    timestamp,
    date_modified   timestamp,
    meal_id         bigint REFERENCES meals (id) ON DELETE CASCADE NOT NULL
);

CREATE TRIGGER set_created
    BEFORE INSERT
    ON meal_food_item_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_created();

CREATE TRIGGER set_modified
    BEFORE UPDATE
    ON meal_food_item_entries
    FOR EACH ROW
EXECUTE PROCEDURE set_date_modified();
