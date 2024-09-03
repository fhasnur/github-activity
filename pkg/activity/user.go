package activity

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type User struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
}

func FetchUserDetails(username string) (User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return User{}, fmt.Errorf("failed to fetch user details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, errors.New("failed to fetch user details, check the username")
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return User{}, fmt.Errorf("error decoding user details: %w", err)
	}

	return user, nil
}

func DisplayUserDetails(user User) {
	color.Cyan("Github User: %s", user.Login)
	if user.Name != "" {
		color.Blue("Name: %s", user.Name)
	}

	if user.Bio != "" {
		color.Yellow("Bio: %s", user.Bio)
	}

	color.Magenta("Public Repos: %d | Followers: %d", user.PublicRepos, user.Followers)
}
