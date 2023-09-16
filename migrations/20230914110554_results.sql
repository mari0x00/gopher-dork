-- +goose Up
-- +goose StatementBegin
CREATE TABLE results(
    id SERIAL PRIMARY KEY,
    config_id INT REFERENCES configs (id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    description TEXT
);
INSERT INTO results(
    config_id, 
    url, 
    description) 
VALUES (
    1,
    'test',  
    'test'), (
    1,
    'test',  
    'test');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
