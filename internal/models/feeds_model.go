package models

import (
	"github.com/google/uuid"
	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/db"
)

func CreateFeed(cfg config.APPConfiger, name, url string, userId uuid.UUID) (*ApiFeed, error) {
	params := db.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: userId,
	}
	feed, err := cfg.GetDB().CreateFeed(cfg.GetCtx(), params)
	if err != nil {
		return nil, err
	}
	apiFeed := databaseFeedToFeed(feed)
	return &apiFeed, nil
}

func GetAllFeeds(cfg config.APPConfiger) ([]ApiFeed, error) {
	feeds, err := cfg.GetDB().GetAllFeeds(cfg.GetCtx())
	if err != nil {
		return nil, err
	}

	return databaseFeedsToFeed(feeds), nil
}

func FollowFeed(cfg config.APPConfiger, feedId, userId uuid.UUID) (*db.FeedsFollow, error) {
	params := db.CreateFeedFollowParams{
		FeedID: feedId,
		UserID: userId,
		ID:     uuid.New(),
	}
	feedFollow, err := cfg.GetDB().CreateFeedFollow(cfg.GetCtx(), params)
	if err != nil {
		return nil, err
	}

	return &feedFollow, nil
}

func GetUserFeeds(cfg config.APPConfiger, userID uuid.UUID) ([]ApiFeed, error) {
	feeds, err := cfg.GetDB().GetUserFeeds(cfg.GetCtx(), userID)
	if err != nil {
		return nil, err
	}

	return databaseFeedsToFeed(feeds), nil
}

func DeleteFeed(cfg config.APPConfiger, id uuid.UUID) error {
	return cfg.GetDB().DeleteFeedFollow(cfg.GetCtx(), id)
}
