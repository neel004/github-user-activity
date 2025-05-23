package commands

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	api "github.com/neel004/github-user-activity/github_api"
)

type EventMessage struct {
	Message     string
	NeedsCount  bool
	NeedsURL    bool
	NeedsAction bool
}

func UserActivity(user string, client *api.Client) error {
	if len(user) <= 0 {
		return fmt.Errorf("user can not be empty.")
	}

	userActivity, err := client.GetUserActivity(user)
	if err != nil {
		fmt.Errorf("%w", err)
	}
	var action string
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Padding(1, 0).
			Foreground(lipgloss.Color("#e979f7")).
			Render(fmt.Sprintf("%s's recent activity(s)", user)),
	)
	for _, activity := range *userActivity {
		switch activity.Type {
		case "ForkEvent":
			action = fmt.Sprintf("forked repo [%s](%s)", activity.Repo.Name, activity.Payload.Forkee.URL)
		case "IssuesEvent":
			action = fmt.Sprintf("%s issue at %s", activity.Payload.Action, activity.Repo.Name)
		case "PullRequestEvent":
			action = fmt.Sprintf("%s pull request at %s", activity.Payload.Action, activity.Repo.Name)
		case "PushEvent":
			action = fmt.Sprintf("pushed %d commits at %s", activity.Payload.Size, activity.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("%s a repository at %s", activity.Payload.Action, activity.Repo.Name)
		}
		actionStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#394e7d")).
			Render(fmt.Sprintf("- %s", action))
		fmt.Println(actionStyle)

	}

	return nil
}
