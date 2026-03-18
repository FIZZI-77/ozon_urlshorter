
-- +goose Up
-- +goose StatementBegin

CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY,
    base_url TEXT  NOT NULL UNIQUE,
    short_url CHAR(10) NOT NULL UNIQUE

);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE urls CASCADE;

-- +goose StatementEnd