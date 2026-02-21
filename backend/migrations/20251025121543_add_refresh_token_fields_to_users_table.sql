-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE users 
    ADD COLUMN refresh_token_hash VARCHAR(255),
    ADD COLUMN refresh_token_expires_at TIMESTAMP WITH TIME ZONE;
    
-- PostgreSQL requires CREATE INDEX to be a separate command
CREATE INDEX idx_users_refresh_token_hash ON users (refresh_token_hash);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE users 
    DROP COLUMN refresh_token_hash,
    DROP COLUMN refresh_token_expires_at;