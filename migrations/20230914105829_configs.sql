-- +goose Up
-- +goose StatementBegin
CREATE TABLE configs(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    query TEXT NOT NULL,
    results_limit INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE configs;
-- +goose StatementEnd
