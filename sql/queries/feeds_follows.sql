-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
RETURNING *;


-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows WHERE id = $1;

-- name: GetUserFeeds :many
SELECT feeds.id, feeds.name, feeds.url, feeds.user_id, feeds.created_at, feeds.updated_at FROM feeds
LEFT JOIN feeds_follows ON feeds.id = feeds_follows.feed_id
LEFT JOIN users ON users.id = feeds_follows.user_id
WHERE users.id = $1;
