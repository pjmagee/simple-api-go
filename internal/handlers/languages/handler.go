// Package languages internal/languages/handler.go
package languages

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Language struct {
	Name                 string  `json:"name"`
	MySkillRating        float32 `json:"stars"`
	ReasonToUse          string  `json:"reason,omitempty"`
	MySkillJustification string  `json:"stars_justification,omitempty"`
}

var languages = []Language{
	{
		Name:                 "Go",
		MySkillRating:        3,
		ReasonToUse:          "Simple and efficient. Popular for microservices and cloud-native applications",
		MySkillJustification: "Currently learning Go due to its simplicity and efficiency",
	},
	{
		Name:                 "Python",
		MySkillRating:        2,
		ReasonToUse:          "Company's preferred language for data engineering and ML",
		MySkillJustification: "Handling ML Models for .NET API delivery, Reviewing Python code used in other teams",
	},
	{
		Name:                 "JavaScript",
		MySkillRating:        2,
		ReasonToUse:          "APIGEE Proxy development, Developer Portal customization",
		MySkillJustification: "Used for APIGEE Proxy development, Developer Portal customization",
	},
	{
		Name:                 "Powershell",
		MySkillRating:        2,
		ReasonToUse:          "Automating tasks and using for simple local scripts",
		MySkillJustification: "Used in combination with KubeCtl, Kiota, and other tools for automating tasks and creating simple local scripts",
	},
	{
		Name:                 "C#",
		MySkillRating:        5,
		ReasonToUse:          "Company's preferred language for API and BFF development",
		MySkillJustification: "Used for API and BFF development, 10 years of experience with the language",
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stars := r.URL.Query().Get("stars")

	var filteredLanguages []Language = make([]Language, 0)
	copy(filteredLanguages, languages)

	if stars != "" {
		for _, language := range languages {
			if rating, err := strconv.ParseFloat(stars, 32); err == nil && float32(rating) == language.MySkillRating {
				filteredLanguages = append(filteredLanguages, language)
			}
		}
	}

	err := json.NewEncoder(w).Encode(filteredLanguages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
