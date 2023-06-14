package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cptvictor95/go-todos/internals/configs"
	"github.com/cptvictor95/go-todos/internals/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		log.Printf("Error loading configs: %v", err)
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.FindMany)
	router.Get("/{id}", handlers.FindOne)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
