-- name: CreatePost :one
insert into posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
values ($1, now(), now(), $2, $3, $4, $5, $6)
returning *;

-- name: GetPosts :many
select *
from posts
join feed_follows on feed_follows.feed_id = posts.feed_id
join users on users.id = feed_follows.user_id
where users.Name = $1;

-- name: GetLatestPosts :many
select *
from posts
join feed_follows on feed_follows.feed_id = posts.feed_id
join users on users.id = feed_follows.user_id
where users.Name = $1
limit $2;
