package routes

import (
	"context"

	"github.com/samvimes01/go-rss/internal/db"
	"github.com/samvimes01/go-rss/internal/env"
)

type APIConfig struct {
	db  *db.Queries
	ctx *context.Context
}

func NewAPIConfig(e *env.Env, ctx *context.Context) *APIConfig {
	return &APIConfig{
		db:  db.InitDb(e),
		ctx: ctx,
	}
}

func (c APIConfig) GetDB() *db.Queries {
	return c.db
}
func (c APIConfig) GetCtx() *context.Context {
	return c.ctx
}
