package tools

import (
	"awesomeProject/internal/usage"
	"encoding/json"
	"net/http"
)

type Tool struct {
	Name   string        `json:"name"`
	Tags   []string      `json:"type"`
	Rating float32       `json:"rating"`
	URL    string        `json:"url"`
	Usage  []usage.Usage `json:"usage"`
}

var tools = []Tool{
	{
		Name:   "GoLand",
		Tags:   []string{"IDE", "Go"},
		Rating: 4,
		URL:    "https://www.jetbrains.com/go/",
		Usage:  []usage.Usage{usage.Past},
	},
	{
		Name:   "Rider",
		Tags:   []string{"IDE", "C#", "Unity"},
		Rating: 4,
		URL:    "https://www.jetbrains.com/rider/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Dagger CICD",
		Tags:   []string{"CI/CD", "DevOps"},
		Rating: 3,
		URL:    "https://dagger.io/",
		Usage:  []usage.Usage{usage.Future},
	},
	{
		Name:   "Visual Studio",
		Tags:   []string{"IDE", "Frontend", "Backend"},
		Rating: 5,
		URL:    "https://visualstudio.microsoft.com/",
		Usage:  []usage.Usage{usage.Past},
	},
	{
		Name:   "Visual Studio Code",
		Tags:   []string{"IDE", "Frontend", "Backend"},
		Rating: 4,
		URL:    "https://code.visualstudio.com/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Postman",
		Tags:   []string{"API", "Testing", "Frontend", "Backend"},
		Rating: 4,
		URL:    "https://www.postman.com/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Git",
		Tags:   []string{"Version Control"},
		Rating: 5,
		URL:    "https://git-scm.com/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Docker",
		Tags:   []string{"Container", "Backend", "DevOps", "Microservices", "Distributed Systems"},
		Rating: 4,
		URL:    "https://www.docker.com/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Kubernetes",
		Tags:   []string{"Container", "Backend", "Orchestration", "Cloud", "DevOps", "Microservices", "Distributed Systems"},
		Rating: 3,
		URL:    "https://kubernetes.io/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Jenkins",
		Tags:   []string{"CI/CD", "DevOps"},
		Rating: 2,
		URL:    "https://www.jenkins.io/",
		Usage:  []usage.Usage{usage.Past},
	},
	{
		Name:   "Microsoft Kiota",
		Tags:   []string{"SDK", "API", "Code Generation", "Multi-Language Support", "OpenAPI"},
		Rating: 5,
		URL:    "https://github.com/microsoft/kiota",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "MongoDB",
		Tags:   []string{"Database", "NoSQL", "Document", "JSON", "BSON"},
		Rating: 5,
		URL:    "https://www.mongodb.com/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "OneDrive",
		Tags:   []string{"Cloud", "Storage", "File Sharing"},
		Rating: 4,
		URL:    "https://www.microsoft.com/en-us/microsoft-365/onedrive/online-cloud-storage",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Nodepad++",
		Tags:   []string{"Text Editor"},
		Rating: 3,
		URL:    "https://notepad-plus-plus.org/",
		Usage:  []usage.Usage{usage.Past},
	},
	{
		Name:   "OneNote",
		Tags:   []string{"Note Taking", "Organization"},
		Rating: 5,
		URL:    "https://www.onenote.com/",
		Usage:  []usage.Usage{usage.Present},
	},
	{
		Name:   "Percipio",
		Tags:   []string{"Learning", "Training", "Skill Development"},
		Rating: 5,
		URL:    "https://www.percipio.com/",
		Usage:  []usage.Usage{usage.Present},
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tools)
}
