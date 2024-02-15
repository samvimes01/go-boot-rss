-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows WHERE id = $1;

-- name: GetUserFeeds :many
SELECT f.id, f.name, f.url, f.user_id, f.created_at, f.updated_at, f.last_fetched_at
FROM feeds AS f
LEFT JOIN feeds_follows ON feeds.id = feeds_follows.feed_id
LEFT JOIN users ON users.id = feeds_follows.user_id
WHERE users.id = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST LIMIT $1 OFFSET $2;

-- name: MarkAsFetched :exec
UPDATE feeds SET last_fetched_at = now(), updated_at = now() WHERE id = $1;