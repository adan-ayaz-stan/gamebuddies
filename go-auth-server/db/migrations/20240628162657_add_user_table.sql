-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    instance_id VARCHAR(255) NOT NULL,
    aud VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    encrypted_password VARCHAR(255) NOT NULL,
    email_confirmed_at TIMESTAMP,
    confirmation_token VARCHAR(255),
    confirmation_sent_at TIMESTAMP,
    recovery_token VARCHAR(255),
    recovery_sent_at TIMESTAMP,
    last_sign_in_at TIMESTAMP,
    banned_until TIMESTAMP
);
CREATE UNIQUE INDEX users_email_idx ON users(email);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE users;