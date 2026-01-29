CREATE TABLE migration_history
(
    id           bigserial PRIMARY KEY,
    name         varchar(128) NOT NULL,
    date_applied timestamp
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
