-- +goose Up
CREATE TABLE IF NOT EXISTS authentication_role
(
    id    TEXT NOT NULL PRIMARY KEY,
    role  TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS authentication_user;