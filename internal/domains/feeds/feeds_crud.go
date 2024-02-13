package feeds

import (
	"github.com/google/uuid"
	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/db"
)

func CreateFeed(cfg config.APPConfiger, name, url string, userId uuid.UUID) (*db.Feed, error) {
	params := db.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: userId,
	}
	user, err := cfg.GetDB().CreateFeed(cfg.GetCtx(), params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllFeeds(cfg config.APPConfiger) ([]db.Feed, error) {
	feeds, err := cfg.GetDB().GetAllFeeds(cfg.GetCtx())
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
