-- +goose Up
-- +goose StatementBegin
-- Images are served from the MinIO public bucket (gm-public).
-- Run `go run ./cmd/seed-games/` from the backend directory after any DB reset
-- to (re-)upload the cover images from frontend/static/games/ to MinIO.
INSERT INTO games (created_at, updated_at, title, subtitle, img_url) 
VALUES 
    (NOW(), NOW(), 'Counter-Strike: Global Offensive', 'Tactical FPS', 'http://localhost:9000/gm-public/csgo.jpg'),
    (NOW(), NOW(), 'Minecraft', 'Sandbox World', 'http://localhost:9000/gm-public/minecraft.jpg'),
    (NOW(), NOW(), 'Valorant', 'Tactical FPS', 'http://localhost:9000/gm-public/valorant.jpg');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM games;
-- +goose StatementEnd
