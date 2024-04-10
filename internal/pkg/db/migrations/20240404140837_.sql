-- +goose Up
-- +goose StatementBegin
ALTER TABLE infsys
    ADD CONSTRAINT fk_resourcesConsumedID FOREIGN KEY (resourcesConsumedID) REFERENCES resources (id);

ALTER TABLE infsys
    ADD CONSTRAINT fk_adminID FOREIGN KEY (adminID) REFERENCES admin (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE infsys
    DROP CONSTRAINT fk_resourcesConsumedID;

ALTER TABLE infsys
    DROP CONSTRAINT fk_adminID;
-- +goose StatementEnd
