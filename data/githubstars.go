package data

import (
	"encoding/json"
	"net/http"

	"git.ran.cafe/maki/foxlib/foxdb"
	"github.com/makinori/blahaj-quest/config"
)

func getGitHubStars() (int, error) {
	req, err := http.NewRequest(
		"GET", "https://api.github.com/repos/"+config.GITHUB_REPO, nil,
	)

	if err != nil {
		return -1, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer res.Body.Close()

	var data struct {
		StargazersCount int `json:"stargazers_count"`
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return -1, err
	}

	return data.StargazersCount, nil
}

var GitHubStars = foxdb.CreateDataCache(
	// start of every hour,
	"githubstars", "0 * * * *", getGitHubStars,
)
