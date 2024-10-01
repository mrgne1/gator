-- name: CreateUser :one
insert into users (id, created_at, updated_at, name)
values ($1, $2, $3, $4)
returning *;

-- name: GetUser :one
select 
    id,
    created_at,
    updated_at,
    name
from users
where name = $1;

-- name: GetUserById :one
select
    id,
    created_at,
    updated_at,
    name
from users
where id = $1;

-- name: GetUsers :many
select
    id,
    created_at,
    updated_at,
    name
from users;

-- name: ResetUserTable :exec
delete from users;

