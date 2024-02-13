package routes

import (
	"encoding/json"
	"net/http"

	"github.com/samvimes01/go-rss/internal/db"
	"github.com/samvimes01/go-rss/internal/domains/feeds"
)

func (cfg *APIConfig) HandleFeedCreate(w http.ResponseWriter, r *http.Request, user *db.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := feeds.CreateFeed(cfg, params.Name, params.Url, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, *feed)
}

func (cfg *APIConfig) HandleFeedGetAll(w http.ResponseWriter, r *http.Request) {
	feeds, err := feeds.GetAllFeeds(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)
}
