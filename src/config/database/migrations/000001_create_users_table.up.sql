CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password BYTEA NOT NULL
);

CREATE INDEX idx_users_id ON users (id);
