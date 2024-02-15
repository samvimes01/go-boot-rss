package routes

import (
	"net/http"

	"github.com/samvimes01/go-rss/internal/api/middleware"
	"github.com/samvimes01/go-rss/internal/env"
)

var apiV = "/v1/"

func makePattern(method string, route string) string {
	return method + " " + apiV + route
}

func Setup(env *env.Env, mux *http.ServeMux, cfg *APIConfig) {
	mux.Handle("/", http.NotFoundHandler())

	mux.HandleFunc(makePattern(http.MethodGet, "readiness"), readyHandler)
	mux.HandleFunc(makePattern(http.MethodGet, "error"), errHandler)

	mux.HandleFunc(makePattern(http.MethodPost, "users"), cfg.HandleUserCreate)
	mux.HandleFunc(makePattern(http.MethodGet, "users"), cfg.HandleUserGetCurrent)

	mux.HandleFunc(makePattern(http.MethodGet, "feeds"), cfg.HandleFeedGetAll)
	mux.HandleFunc(makePattern(http.MethodPost, "feeds"), middleware.Auth(cfg, cfg.HandleFeedCreate))

	mux.HandleFunc(makePattern(http.MethodGet, "feed_follows"), middleware.Auth(cfg, cfg.HandleFeedFollowsGetMany))
	mux.HandleFunc(makePattern(http.MethodPost, "feed_follows"), middleware.Auth(cfg, cfg.HandleFeedFollowsCreate))
	mux.HandleFunc(makePattern(http.MethodDelete, "feed_follows/{feedFollowID}"), middleware.Auth(cfg, cfg.HandleFeedFollowsDelete))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, resp{"ok"})
}
func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
