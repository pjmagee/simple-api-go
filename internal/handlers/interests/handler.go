package interests

import (
	"encoding/json"
	"net/http"
)

type Interest struct {
	Name string   `json:"name"`
	Tags []string `json:"type"`
}

var interests = []Interest{
	{
		Name: "Reading Coding books",
		Tags: []string{"Book", "Coding"},
	},
	{
		Name: "Games",
		Tags: []string{"PC", "XBox"},
	},
	{
		Name: "Dune Series",
		Tags: []string{"Book", "Sci-Fi"},
	},
	{
		Name: "Star Wars",
		Tags: []string{"Movie", "Sci-Fi"},
	},
	{
		Name: "Star Trek",
		Tags: []string{"Series", "Sci-Fi"},
	},
	{
		Name: "The Expanse",
		Tags: []string{"Series", "Sci-Fi"},
	},
	{
		Name: "The Mandalorian",
		Tags: []string{"Series", "Sci-Fi", "Star Wars"},
	},
	{
		Name: "Guinness",
		Tags: []string{"Beer", "Stout"},
	},
	{
		Name: "Jake Gyllenhaal",
		Tags: []string{"Actor", "Movies"},
	},
	{
		Name: "Tom Hanks",
		Tags: []string{"Actor", "Movies"},
	},
	{
		Name: "Keanu Reeves",
		Tags: []string{"Actor", "Movies"},
	},
	{
		"JetBrains",
		[]string{"Tool", "IDE"},
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(interests)
}
