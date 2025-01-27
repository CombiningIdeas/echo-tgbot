-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_links (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(256) NOT NULL,
    name_link VARCHAR(4096) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_links;
-- +goose StatementEnd
