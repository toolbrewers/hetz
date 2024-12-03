CREATE TABLE IF NOT EXISTS sessions (
    id              INTEGER     PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER     NOT NULL,
    token           TEXT        NOT NULL,
    expires_at      DATETIME    NOT NULL,
    created_at      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      DATETIME,
    deleted_at      DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
