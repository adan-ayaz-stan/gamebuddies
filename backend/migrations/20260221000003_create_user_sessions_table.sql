-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_sessions (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,

    user_id    BIGINT    NOT NULL,
    started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    ended_at   TIMESTAMP WITH TIME ZONE NULL,      -- NULL means currently active

    CONSTRAINT fk_user_sessions_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX idx_user_sessions_user_id   ON user_sessions (user_id);
CREATE INDEX idx_user_sessions_deleted_at ON user_sessions (deleted_at);
CREATE INDEX idx_user_sessions_ended_at  ON user_sessions (ended_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_sessions;
-- +goose StatementEnd
