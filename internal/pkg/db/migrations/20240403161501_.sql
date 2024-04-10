-- +goose Up
-- +goose StatementBegin
create table resources
(
    id  BIGINT PRIMARY KEY NOT NULL,
    cpu INTEGER            NOT NULL DEFAULT 0,
    ram INTEGER            NOT NULL DEFAULT 0,
    hdd TEXT               NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table resources;
-- +goose StatementEnd
