package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cptvictor95/go-todos/internals/models"
)

func FindMany(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error trying to find todos: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
