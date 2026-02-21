-- +goose Up
-- +goose StatementBegin
ALTER TABLE profiles ADD COLUMN IF NOT EXISTS total_matches INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE profiles DROP COLUMN IF EXISTS total_matches;
-- +goose StatementEnd
