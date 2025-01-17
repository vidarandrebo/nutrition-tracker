CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(128),
    passwordhash BYTEA
);

CREATE TABLE IF NOT EXISTS fooditems (
    id INTEGER PRIMARY KEY,
    name VARCHAR(128)
);

insert into fooditems(id, name) values (1, 'hello');