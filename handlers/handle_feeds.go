package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yigiterdev/rss-aggregator/internal/database"
	"github.com/yigiterdev/rss-aggregator/models"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create user: %v", err))
		return
	}

	respondWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot get feeds: %v", err))
		return
	}
	respondWithJSON(w, 200, models.DatabaseFeedsToFeeds(feeds))
}
