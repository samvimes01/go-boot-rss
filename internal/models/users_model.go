package models

import (
	"github.com/google/uuid"
	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/db"
)

func CreateUser(cfg config.APPConfiger, name string) (*db.User, error) {
	params := db.CreateUserParams{
		Name: name,
		ID:   uuid.New(),
	}
	ctx := cfg.GetCtx()
	user, err := cfg.GetDB().CreateUser(*ctx, params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByApiKey(cfg config.APPConfiger, key string) (*db.User, error) {
	ctx := cfg.GetCtx()
	user, err := cfg.GetDB().GetUserByApiKey(*ctx, key)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
