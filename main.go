package main

import (
	"log"
	"net/http"
	"targeting-engine/db"
	"targeting-engine/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	db.InitDB()

	r := chi.NewRouter()
	r.Get("/v1/delivery", handlers.DeliveryHandler)

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
