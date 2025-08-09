package recipes

import "time"

type Recipe struct {
	ID          any       `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Slug        string    `json:"slug" bson:"slug"`
	ImageURL    string    `json:"image_url" bson:"image_url"`
	Description []string  `json:"description" bson:"description"`
	Time        string    `json:"time" bson:"time"`
	Ingredients []string  `json:"ingredients" bson:"ingredients"`
	Steps       []string  `json:"steps" bson:"steps"`
	Published   bool      `json:"published" bson:"published"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
