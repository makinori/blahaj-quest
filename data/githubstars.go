package data

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/makinori/blahaj-quest/config"
)

func GetGitHubStars() int {
	var cache int
	err := GetCache("githubStars", &cache)
	if err == nil {
		return cache
	}

	req, err := http.NewRequest(
		"GET", "https://api.github.com/repos/"+config.GitHubRepo, nil,
	)

	if err != nil {
		slog.Error("failed to make req for github stars", "err", err)
		return -1
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		slog.Error("failed to get github stars", "err", err)
		return -1
	}
	defer res.Body.Close()

	var data struct {
		StargazersCount int `json:"stargazers_count"`
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		slog.Error("failed to decode github stars", "err", err)
		return -1
	}

	SetCache("githubStars", data.StargazersCount)

	return data.StargazersCount
}
