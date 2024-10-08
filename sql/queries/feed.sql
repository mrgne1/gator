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

-- name: DeleteFeedFollows :exec
delete from feed_follows
where feed_follows.user_id = (select id from users where users.name = $1)
    and feed_follows.feed_id = (select id from feeds where feeds.url = $2);

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

-- name: MarkFeedFetched :exec
update feeds
set last_fetched_at = now(),
    updated_at = now()
where id = $1;

-- name: GetNextFeedToFetch :one
select 
    id,
    name,
    url,
    user_id,
    created_at,
    updated_at
from feeds
order by last_fetched_at asc nulls first
limit 1;

-- name: ResetFeedTable :exec
delete from feeds;

