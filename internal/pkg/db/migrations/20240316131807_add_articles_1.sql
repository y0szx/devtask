-- +goose Up
-- +goose StatementBegin
create table listinfsys
(
    id       BIGSERIAL PRIMARY KEY NOT NULL,
    name     TEXT                  NOT NULL DEFAULT '',
    owner    TEXT                  NOT NULL DEFAULT '',
    admin    TEXT                  NOT NULL DEFAULT '',
    contacts TEXT                  NOT NULL DEFAULT '',
    infsysID BIGINT REFERENCES infsys (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table listinfsys;
-- +goose StatementEnd
