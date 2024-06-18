-- +goose Up
-- +goose StatementBegin
create table images
(
    image_id   BIGSERIAL PRIMARY KEY NOT NULL,
    id         BIGINT                NOT NULL,
    image_data BYTEA                 NOT NULL,
    image_name TEXT                  NOT NULL,
    FOREIGN KEY (id) REFERENCES listinfsys (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table images;
-- +goose StatementEnd
