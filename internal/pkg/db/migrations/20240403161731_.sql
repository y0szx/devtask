-- +goose Up
-- +goose StatementBegin
create table admin
(
    id       BIGINT PRIMARY KEY NOT NULL,
    fio      TEXT               NOT NULL DEFAULT '',
    email    TEXT               NOT NULL DEFAULT '',
    telegram TEXT               NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table admin;
-- +goose StatementEnd
