CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

drop table if exists users cascade;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name character varying(255) NOT NULL,
                       last_name character varying(255) NOT NULL,
                       user_active integer NOT NULL DEFAULT 0,
                       email character varying(255) NOT NULL UNIQUE,
                       password character varying(60) NOT NULL,
                       created_at timestamp without time zone NOT NULL DEFAULT now(),
                       updated_at timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


drop table if exists tokens_jwt;

CREATE TABLE tokens_jwt (
                        id SERIAL PRIMARY KEY,
                        user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
                        first_name character varying(255) NOT NULL,
                        email character varying(255) NOT NULL,
                        token character varying(255) NOT NULL,
                        created_at timestamp without time zone NOT NULL DEFAULT now(),
                        updated_at timestamp without time zone NOT NULL DEFAULT now(),
                        expiry timestamp without time zone NOT NULL
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON tokens_jwt
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();