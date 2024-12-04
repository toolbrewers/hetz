CREATE TABLE IF NOT EXISTS roles (
    name            TEXT        NOT NULL PRIMARY KEY,
    created_at      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      DATETIME,
    deleted_at      DATETIME
);

CREATE TABLE IF NOT EXISTS sessions (
    id              INTEGER     PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER     NOT NULL,
    role_id         TEXT        NOT NULL,
    token           TEXT        NOT NULL, -- Varchar(64) is all that is needed
    expires_at      DATETIME    NOT NULL,
    created_at      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      DATETIME,
    deleted_at      DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(name),
    INDEX (token)
);
