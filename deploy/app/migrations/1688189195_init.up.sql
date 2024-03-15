CREATE TABLE IF NOT EXISTS users
(
    id         BIGINT                                    NOT NULL PRIMARY KEY,
    first_name VARCHAR(64)                               NOT NULL,
    last_name  VARCHAR(64)                               NOT NULL,
    username   VARCHAR(64)                               NOT NULL,
    language   CHAR(3)                                   NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE               NULL
);

ALTER TABLE users
    OWNER TO exifbot;
