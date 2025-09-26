CREATE TABLE migration_history
(
    id           bigserial PRIMARY KEY,
    name         varchar(128) NOT NULL,
    date_applied timestamp
);
