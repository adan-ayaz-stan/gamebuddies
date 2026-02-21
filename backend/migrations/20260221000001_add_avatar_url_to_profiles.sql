-- +goose Up
-- +goose StatementBegin
ALTER TABLE profiles ADD COLUMN IF NOT EXISTS avatar_url VARCHAR(512) DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE profiles DROP COLUMN IF EXISTS avatar_url;
-- +goose StatementEnd
