-- +goose Up
-- +goose StatementBegin
CREATE TABLE results(
    id SERIAL PRIMARY KEY,
    config_id INT REFERENCES configs (id) ON DELETE CASCADE,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    status INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
