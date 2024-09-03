package activity

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
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
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

func FetchGithubActivity(username string) error {
	user, err := FetchUserDetails(username)
	if err != nil {
		return err
	}
	DisplayUserDetails(user)

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return errors.New("user not found, please check the username")
		}
		return fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	var events []GithubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found for this user")
		return nil
	}

	color.Cyan("\nRecent Activities:")
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
		color.Green("- %s (at %s)", action, event.CreatedAt.Format(time.RFC3339))
		fmt.Println("-------------------------------------------------")
	}

	return nil
}
