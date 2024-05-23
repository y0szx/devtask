-- +goose Up
-- +goose StatementBegin
create table infsys
(
    id                 BIGINT PRIMARY KEY NOT NULL,
    name               TEXT               NOT NULL DEFAULT '',
    owner              TEXT               NOT NULL DEFAULT '',
    vms                TEXT               NOT NULL DEFAULT '',
    cpu                BIGINT             NOT NULL DEFAULT 0,
    ram                BIGINT             NOT NULL DEFAULT 0,
    hdd                TEXT               NOT NULL DEFAULT '',
    software_used       TEXT               NOT NULL DEFAULT '',
    admin_name          TEXT               NOT NULL DEFAULT '',
    admin_email         TEXT               NOT NULL DEFAULT '',
    admin_tg            TEXT               NOT NULL DEFAULT '',
    resource_assignment TEXT               NOT NULL DEFAULT '',
    status             BOOLEAN            NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table infsys;
-- +goose StatementEnd
