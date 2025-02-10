CREATE TABLE IF NOT EXISTS users
(
    id            bigserial primary key,
    name          varchar(128),
    email         varchar(128),
    password_hash bytea
);

CREATE TABLE IF NOT EXISTS food_items
(
    id           bigserial primary key,
    manufacturer varchar(128),
    product      varchar(128),
    owner        bigint REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS macronutrients
(
    id           bigserial primary key,
    protein      double precision,
    carbohydrate double precision,
    fat          double precision,
    kcal         double precision
);

CREATE TABLE IF NOT EXISTS micronutrients
(
    id     bigserial primary key,
    name   varchar,
    amount double precision
);


CREATE TABLE IF NOT EXISTS days
(
    id   bigserial primary key,
    date date
);


CREATE TABLE IF NOT EXISTS meals
(
    id              bigserial primary key,
    sequence_number integer,
    day_id          bigint REFERENCES days (id)
);

CREATE TABLE IF NOT EXISTS food_items_macronutrients
(
    food_item_id     bigint references food_items (id),
    macronutrient_id bigint references macronutrients (id),
    primary key (food_item_id, macronutrient_id)
);

CREATE TABLE IF NOT EXISTS food_items_micronutrients
(
    food_item_id     bigint references food_items (id),
    micronutrient_id bigint references micronutrients (id),
    primary key (food_item_id, micronutrient_id)
);
