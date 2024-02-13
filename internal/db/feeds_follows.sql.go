// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feeds_follows.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
RETURNING id, user_id, feed_id, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow, arg.ID, arg.UserID, arg.FeedID)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows WHERE id = $1
`

func (q *Queries) DeleteFeedFollow(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, id)
	return err
}

const getUserFeeds = `-- name: GetUserFeeds :many
SELECT feeds.id, feeds.name, feeds.url, feeds.user_id, feeds.created_at, feeds.updated_at FROM feeds
LEFT JOIN feeds_follows ON feeds.id = feeds_follows.feed_id
LEFT JOIN users ON users.id = feeds_follows.user_id
WHERE users.id = $1
`

func (q *Queries) GetUserFeeds(ctx context.Context, id uuid.UUID) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getUserFeeds, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}