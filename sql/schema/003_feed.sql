-- +goose Up
create table feeds (
    id uuid primary key,
    name text not null,
    url text unique not null,
    user_id uuid references users (id) on delete cascade not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table feeds;
