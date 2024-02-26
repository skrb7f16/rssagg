package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/skrb7f16/rssagg/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, 500, "Something went wrong")
		return
	}
	feeds_follow, err := apiCfg.DB.CreateFeedsFollow(r.Context(), database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		log.Printf("%v", err)
		responseWithErr(w, 400, "Please provide name %v")
	}
	responseWithJson(w, 201, convertDbFeedsFollowToNormalFeedsFollow(feeds_follow))
}
