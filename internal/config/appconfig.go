package config

import (
	"context"

	"github.com/samvimes01/go-rss/internal/db"
)

type APPConfiger interface {
	GetDB() *db.Queries
	GetCtx() *context.Context
}
