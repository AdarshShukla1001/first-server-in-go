-- +goose Up


CREATE TABLE users (
    id UUID TIMESTAMP NOT NULL
    created_at TIMESTAMP NOT NULL
    updated_at TIMESTAMP NOT NULL
    name TEXT NOT NULL
);

-- +goose Down

