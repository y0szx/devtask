-- +goose Up
-- +goose StatementBegin
create table infsys
(
    id                  BIGINT PRIMARY KEY NOT NULL,
    name                TEXT               NOT NULL DEFAULT '',
    owner               TEXT               NOT NULL DEFAULT '',
    vms                 TEXT               NOT NULL DEFAULT '',
    softwareUsed        TEXT               NOT NULL DEFAULT '',
    resourceAssignment  TEXT               NOT NULL DEFAULT '',
    status              BOOLEAN            NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table infsys;
-- +goose StatementEnd
