package middlewares

import (
	"net/http"

	"github.com/samvimes01/go-rss/internal/auth"
	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/db"
	"github.com/samvimes01/go-rss/internal/domains/users"
)

type authedHandler func(http.ResponseWriter, *http.Request, *db.User)

func Auth(cfg config.APPConfiger, handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if user, err := isAuthorized(cfg, r); err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		} else {
			handler(w, r, user)
		}
	}
}

func isAuthorized(cfg config.APPConfiger, r *http.Request) (*db.User, error) {
	key, err := auth.ParseApiKeyHeader(r)
	if err != nil {
		return &db.User{}, err
	}

	user, err := users.GetUserByApiKey(cfg, key)
	if err != nil {
		return &db.User{}, err
	}
	return user, nil
}
