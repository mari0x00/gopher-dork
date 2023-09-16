-- +goose Up
-- +goose StatementBegin
CREATE TABLE configs(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    query TEXT NOT NULL,
    results_limit INT NOT NULL,
    proxy TEXT
);
INSERT INTO configs(
    name, 
    query, 
    results_limit,
    proxy) 
VALUES (
    'test', 
    'ext:(doc | docx | pdf | xls | xlsx | txt | ps | rtf | odt | sxw | psw | ppt | pps | xml | ppt | pptx) (intext:"Internal" | intext:"Confidential" | intext:"Confidential - employees only")', 
    '5',
    'none');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE configs;
-- +goose StatementEnd
