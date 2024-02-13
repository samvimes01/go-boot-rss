package routes

import (
	"encoding/json"
	"net/http"

	"github.com/samvimes01/go-rss/internal/auth"
	"github.com/samvimes01/go-rss/internal/models"
)

func (cfg *APIConfig) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := models.CreateUser(cfg, params.Name)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, *user)
}

func (cfg *APIConfig) HandleUserGetCurrent(w http.ResponseWriter, r *http.Request) {
	key, err := auth.ParseApiKeyHeader(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := models.GetUserByApiKey(cfg, key)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find user")
		return
	}

	respondWithJSON(w, http.StatusCreated, *user)
}
