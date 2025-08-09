package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/sindriax/moms-recipes-api/internal/db"
	"github.com/sindriax/moms-recipes-api/internal/recipes"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })

	col := db.GetCollection("recipes")
	h := &recipes.Handler{Col: col}
	r.Get("/recipes", h.List)
	r.Post("/recipes", h.Create)

	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
