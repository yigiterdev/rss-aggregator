package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yigiterdev/rss-aggregator/internal/auth"
	"github.com/yigiterdev/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(newUser))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Auth error: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}
