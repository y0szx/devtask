-- +goose Up
-- +goose StatementBegin
create table docs
(
    id       BIGINT NOT NULL,
    doc_data BYTEA  NOT NULL,
    doc_name TEXT   NOT NULL,
    FOREIGN KEY (id) REFERENCES listinfsys (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table docs;
-- +goose StatementEnd
