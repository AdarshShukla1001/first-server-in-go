package main

import (
	"encoding/json"
	"fmt"


	"net/http"
	"time"

	"github.com/AdarshShukla1001/first-go-server/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)






func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}




func (cfg *apiConfig) handlerFeedFollowsDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	fellowFollowIDStr:= chi.URLParam(r,"feedFollowID")
	feedFollowID,err:= uuid.Parse(fellowFollowIDStr)

	if err!=nil {
		respondWithError(w,400,fmt.Sprintf("Could not parse feed follow ID: %v",err))
	}
	
	
	feedFollowDeleteErr := cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if feedFollowDeleteErr != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not delete feed follow: %v",err))
		return
	}

	respondWithJson(w, http.StatusOK,struct{}{})
}




func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Couldn't create feed %v", err))
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}
