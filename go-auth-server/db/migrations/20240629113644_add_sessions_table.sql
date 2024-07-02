-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE sessions (
    id INT AUTO_INCREMENT NOT NULL,
    user_id CHAR(36) NOT NULL,
    not_after TIMESTAMP,
    refreshed_at TIMESTAMP,
    user_agent VARCHAR(255) NOT NULL,
    ip VARCHAR(255),
    tag VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
ALTER TABLE sessions
ADD FOREIGN KEY (user_id) REFERENCES users(id);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE sessions;