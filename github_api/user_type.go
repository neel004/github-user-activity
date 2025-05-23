package github_api

type UserActivity []struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Forkee struct {
			URL string `json:"html_url"`
		} `json:"forkee"`
		Action string `json:"action"`
		Issue  struct {
			URL string `json:"url"`
		}
		PullRequest struct {
			URL string `json:"url"`
		}
		Size int `json:"size"`
	} `json:"payload"`
}
