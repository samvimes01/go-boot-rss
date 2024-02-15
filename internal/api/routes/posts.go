package routes

import (
	"net/http"
	"strconv"

	"github.com/samvimes01/go-rss/internal/db"
)

func (cfg *APIConfig) HandlPostsGetForUser(w http.ResponseWriter, r *http.Request, user *db.User) {
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	ctx := cfg.GetCtx()
	params := db.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
		Offset: 0,
	}
	posts, err := cfg.GetDB().GetPostsByUser(*ctx, params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
