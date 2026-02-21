-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE profiles (
    -- gorm.Model fields
    id BIGSERIAL PRIMARY KEY, -- PostgreSQL auto-incrementing primary key
    created_at TIMESTAMP with time zone NOT NULL,
    updated_at TIMESTAMP with time zone NOT NULL,
    deleted_at TIMESTAMP with time zone NULL,

    -- Profile struct fields
    user_id BIGINT NOT NULL UNIQUE,

    display_name VARCHAR(255) NOT NULL,

    -- Define the Foreign Key constraint
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

-- Create an index on the user_id foreign key
CREATE INDEX idx_profiles_user_id ON profiles (user_id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE profiles;