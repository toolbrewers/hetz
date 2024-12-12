CREATE TABLE IF NOT EXISTS sessions (
    id              INTEGER     PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER     NOT NULL,
    token           TEXT        UNIQUE NOT NULL, -- Varchar(64) is all that is needed
    expires_at      DATETIME    NOT NULL,
    user_agent      TEXT        NOT NULL,
    ip_address      TEXT        NOT NULL,
    created_at      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      DATETIME,
    deleted_at      DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_token ON sessions(token);

CREATE TABLE IF NOT EXISTS roles (
    id              INTEGER     PRIMARY KEY AUTOINCREMENT,
    name            TEXT        NOT NULL,
    created_at      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      DATETIME,
    deleted_at      DATETIME
);