// Package actors internal/handlers/actors/handler.go
package actors

import (
	"encoding/json"
	"net/http"
)

type Actor struct {
	Name    string  `json:"name"`
	ImdbUrl string  `json:"imdb_url"`
	Stars   float32 `json:"stars"`
}

var actors = []Actor{
	{
		Name:    "Jake Gyllenhaal",
		ImdbUrl: "https://www.imdb.com/name/nm0350453/",
		Stars:   5,
	},
	{
		Name:    "Tom Hanks",
		ImdbUrl: "https://www.imdb.com/name/nm0000158/",
		Stars:   5,
	},
	{
		Name:    "Keanu Reeves",
		ImdbUrl: "https://www.imdb.com/name/nm0000206/",
		Stars:   5,
	},
	{
		Name:    "Michael C. Hall",
		ImdbUrl: "https://www.imdb.com/name/nm0350453/",
		Stars:   5,
	},
	{
		Name:    "Mads Mikkelsen",
		ImdbUrl: "https://www.imdb.com/name/nm0586568/",
		Stars:   5,
	},
	{
		Name:    "Sigourney Weaver",
		ImdbUrl: "https://www.imdb.com/name/nm0000244/",
		Stars:   5,
	},
	{
		Name:    "Natalie Portman",
		ImdbUrl: "https://www.imdb.com/name/nm0000204/",
		Stars:   5,
	},
	{
		Name:    "Antony Hopkins",
		ImdbUrl: "https://www.imdb.com/name/nm0000164/",
		Stars:   5,
	},
	{
		Name:    "Morgan Freeman",
		ImdbUrl: "https://www.imdb.com/name/nm0000151/",
		Stars:   5,
	},
	{
		Name:    "Denzel Washington",
		ImdbUrl: "https://www.imdb.com/name/nm0000243/",
		Stars:   5,
	},
	{
		Name:    "Robert De Niro",
		ImdbUrl: "https://www.imdb.com/name/nm0000134/",
		Stars:   5,
	},
	{
		Name:    "Leonardo DiCaprio",
		ImdbUrl: "https://www.imdb.com/name/nm0000138/",
		Stars:   5,
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actors)
}
