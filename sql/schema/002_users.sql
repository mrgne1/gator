-- +goose Up
alter table users
add constraint users_name_uq unique(name);

-- +goose Down
alter table users
drop constraint users_name_uq;
