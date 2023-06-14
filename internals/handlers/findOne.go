package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cptvictor95/go-todos/internals/models"
	"github.com/go-chi/chi/v5"
)

func FindOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error parsing id: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	var todo models.Todo

	todo, err = models.Get(int64(id))

	if err != nil {
		log.Printf("Error trying to find todo: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
