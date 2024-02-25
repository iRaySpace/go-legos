-- +goose Up
CREATE TABLE users (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username    TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
