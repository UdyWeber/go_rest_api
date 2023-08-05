package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"udyweber/go_rest_api/models"
	"udyweber/go_rest_api/router"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Method Not Allowed")
		return
	}

	var person models.User
	state := router.GlobalState

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't decode JSON: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	newHash := uuid.New().String()
	
	person.Id = newHash
	state.Set(newHash, person)

	jsonData, err := json.Marshal(person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Method Not Allowed")
		return
	}

	users := router.GlobalState.GetAll()
	
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}