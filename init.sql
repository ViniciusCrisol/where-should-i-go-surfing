CREATE TABLE users
(
    id         VARCHAR(36)  NOT NULL UNIQUE PRIMARY KEY,
    name       VARCHAR(256) NOT NULL,
    email      VARCHAR(256) NOT NULL UNIQUE,
    password   VARCHAR(256) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at TIMESTAMP    NOT NULL DEFAULT now()
);
CREATE INDEX index_users_email ON users (email);
