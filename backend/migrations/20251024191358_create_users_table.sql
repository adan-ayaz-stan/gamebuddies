-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE users (
    -- gorm.Model fields
    id BIGSERIAL PRIMARY KEY,  -- PostgreSQL auto-incrementing primary key
    created_at TIMESTAMP with time zone NOT NULL,
    updated_at TIMESTAMP with time zone NOT NULL,
    deleted_at TIMESTAMP with time zone NULL,

    -- User struct fields
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Create an index on deleted_at for soft deletion queries
CREATE INDEX idx_users_deleted_at ON users (deleted_at);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE users;