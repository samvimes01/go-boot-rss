package routes

import (
	"net/http"

	"github.com/samvimes01/go-rss/internal/env"
)

var apiV = "/v1/"

func makePattern(method, route string) string {
	return method + " " + apiV + route
}

func Setup(env *env.Env, mux *http.ServeMux, cfg *APIConfig) {
	mux.HandleFunc(makePattern(http.MethodGet, "readiness"), readyHandler)
	mux.HandleFunc(makePattern(http.MethodGet, "error"), errHandler)

	mux.HandleFunc(makePattern(http.MethodPost, "users"), cfg.HandleUserCreate)
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
