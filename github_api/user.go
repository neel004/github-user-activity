package github_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (c *Client) GetUserActivity(user string) (*UserActivity, error) {
	url := fmt.Sprintf(baseURL + "users/" + user + "/events/public")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error encountered while creating request.")
	}

	token, err_bool := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	if !err_bool {
		return nil, fmt.Errorf("please add GITHUB_ACCESS_TOKEN in the .env file.")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer "+token))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error encountered while reading response.")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error encountered while making request. Status: %s", resp.Status)
	}

	defer resp.Body.Close()
	userActivity := UserActivity{}

	err = json.NewDecoder(resp.Body).Decode(&userActivity)
	if err != nil {
		return nil, fmt.Errorf("error encountered while reading/parsing data.")
	}
	return &userActivity, nil
}
