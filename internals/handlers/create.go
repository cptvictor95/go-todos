package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cptvictor95/go-todos/internals/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error decoding request body to json: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	id, err := models.Insert(todo)

	var res map[string]any

	// How do I make this response handling better based on status codes?
	if err != nil {
		res = map[string]any{
			"Error": true,
			// How do I do this with log instead of fmt?
			"Message": fmt.Sprintf("An error has occurred when inserting a todo: %v", err),
		}
	} else {
		res = map[string]any{
			"Error": false,
			// How do I do this with log instead of fmt?
			"Message": fmt.Sprintf("Todo inserted successfully! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}
