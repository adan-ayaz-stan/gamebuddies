-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id CHAR(36),
    instance_id VARCHAR(255) DEFAULT "000-000-000",
    token VARCHAR(255),
    session_id INT NOT NULL,
    revoked BOOLEAN DEFAULT false,
    parent VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
ALTER TABLE refresh_tokens
ADD FOREIGN KEY (session_id) REFERENCES sessions(id);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE refresh_tokens;