package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/samvimes01/go-rss/internal/db"
)

type ApiFeed struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Url           string    `json:"url"`
	UserID        uuid.UUID `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
}

func databaseFeedsToFeed(feed []db.Feed) []ApiFeed {
	var feeds []ApiFeed
	for _, f := range feed {
		feeds = append(feeds, databaseFeedToFeed(f))
	}

	return feeds
}

func databaseFeedToFeed(feed db.Feed) ApiFeed {
	var lastFetchedAt time.Time
	if feed.LastFetchedAt.Valid {
		lastFetchedAt = feed.LastFetchedAt.Time
	}

	return ApiFeed{
		ID:            feed.ID,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		LastFetchedAt: lastFetchedAt,
	}
}
