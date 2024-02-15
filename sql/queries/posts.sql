-- name: CreatePost :one
INSERT INTO posts (title, description, url, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsByUser :many
SELECT * FROM posts p
LEFT JOIN feeds f ON p.feed_id = f.id
WHERE f.user_id = $1
ORDER BY published_at DESC
LIMIT $2 OFFSET $3;
