-- name: CreateFeed :one 
insert into feeds (id, name, url, user_id, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select * from feeds;

-- name: CreateFeedFollows :one
insert into feed_follows (id, user_id, feed_id, created_at, updated_at)
values ($1, $2, $3, $4, $5)
returning id,
    user_id,
    feed_id,
    created_at,
    updated_at,
    (select name from users where users.id = user_id) as user,
    (select name from feeds where feeds.id = feed_id) as feed;

-- name: GetFeedByUrl :one
select
    id,
    name,
    url,
    user_id,
    created_at,
    updated_at
from feeds
where url = $1;

-- name: GetFeedFollowsForUser :many
select
    feeds.name as feed,
    users.name as user
from users
join feed_follows on feed_follows.user_id = users.id
join feeds on feeds.id = feed_follows.feed_id
where users.name = $1;

-- name: ResetFeedTable :exec
delete from feeds;

