-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE games (
    id BIGSERIAL PRIMARY KEY,  -- PostgreSQL auto-incrementing primary key
    created_at TIMESTAMP with time zone NOT NULL,
    updated_at TIMESTAMP with time zone NOT NULL,
    deleted_at TIMESTAMP with time zone NULL,

    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(255),
    img_url VARCHAR(255) NOT NULL
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE games;