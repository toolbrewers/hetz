-- Create users table schema
CREATE TABLE users (
    id              BIGSERIAL       NOT NULL,
    username        VARCHAR(32)     UNIQUE NOT NULL,
    email           VARCHAR(255)    UNIQUE NOT NULL,
    password        TEXT            NOT NULL,
    hetzner_token   TEXT            UNIQUE NOT NULL,
    created_at      TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP       NULL,
    deleted_at      TIMESTAMP       NULL,
    PRIMARY KEY(id)
);
