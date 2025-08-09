package recipes

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Col *mongo.Collection
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	cur, err := h.Col.Find(r.Context(), bson.M{"published": true})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var out []Recipe
	if err := cur.All(r.Context(), &out); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var in Recipe
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad json", 400)
		return
	}
	now := time.Now()
	in.CreatedAt = now
	if !in.Published {
		in.Published = true
	}
	in.Slug = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(in.Name), " ", "-"))

	res, err := h.Col.InsertOne(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(res)
}
