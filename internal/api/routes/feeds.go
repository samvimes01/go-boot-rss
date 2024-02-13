package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/samvimes01/go-rss/internal/db"
	"github.com/samvimes01/go-rss/internal/models"
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

	feed, err := models.CreateFeed(cfg, params.Name, params.Url, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	feedFollow, err := models.FollowFeed(cfg, feed.ID, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow feed")
		return
	}

	resp := struct {
		Feed       *models.ApiFeed        `json:"feed"`
		FeedFollow *db.FeedsFollow `json:"feed_follow"`
	}{feed, feedFollow}

	respondWithJSON(w, http.StatusCreated, resp)
}

func (cfg *APIConfig) HandleFeedGetAll(w http.ResponseWriter, r *http.Request) {
	feeds, err := models.GetAllFeeds(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *APIConfig) HandleFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user *db.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedID, err := uuid.Parse(params.FeedID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode feed_id")
		return
	}

	feedFollow, err := models.FollowFeed(cfg, feedID, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFollow)
}

func (cfg *APIConfig) HandleFeedFollowsGetMany(w http.ResponseWriter, r *http.Request, user *db.User) {
	feeds, err := models.GetAllFeeds(cfg)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)
}

func (cfg *APIConfig) HandleFeedFollowsDelete(w http.ResponseWriter, r *http.Request, user *db.User) {
	feedFollowID := r.URL.Query().Get("feedFollowID")
	id, err := uuid.Parse(feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode feedFollowID")
		return
	}

	err = models.DeleteFeed(cfg, id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feeds")
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
