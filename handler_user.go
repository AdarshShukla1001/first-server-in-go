package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AdarshShukla1001/first-go-server/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramerters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramerters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing Json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create user : %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w,200,databaseUserToUser(user))
}
