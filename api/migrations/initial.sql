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
    protein      double precision,
    carbohydrate double precision,
    fat          double precision,
    kcal         double precision,
    source       varchar(128),
    owner_id     bigint REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS micronutrients
(
    id           bigserial primary key,
    name         varchar,
    amount       double precision,
    food_item_id bigint REFERENCES food_items (id)
);


CREATE TABLE IF NOT EXISTS meals
(
    id              bigserial primary key,
    sequence_number integer,
    meal_time       timestamp,
    owner_id        bigint REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS meal_entries
(
    id           bigserial primary key,
    amount       double precision,
    food_item_id bigint references food_items (id),
    meal_id      bigint REFERENCES meals (id)
)
