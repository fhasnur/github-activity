package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	Name string `json:"name"`
}

type Author struct {
	Name string `json:"name"`
}

type Commit struct {
	Message string `json:"message"`
}

type Payload struct {
	Commits []Commit `json:"commits"`
	Action  string   `json:"action"`
	RefType string   `json:"ref_type"`
}

type GithubEvent struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

func FetchGithubActivity(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data", err)
		return
	}
	defer resp.Body.Close()

	var events []GithubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}

	for _, event := range events {
		var action string
		switch event.Type {
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commit(s) to %s\n", commitCount, event.Repo.Name)
		case "IssueEvent":
			action = fmt.Sprintf("%s an issue in %s", event.Payload.Action, event.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", event.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", event.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("Created %s in %s", event.Payload.RefType, event.Repo.Name)
		default:
			action = fmt.Sprintf("%s in %s", event.Type, event.Repo.Name)
		}
		fmt.Println("- ", action)
	}
}
