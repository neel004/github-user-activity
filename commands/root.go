package commands

import (
	"fmt"
	api "github.com/neel004/github-user-activity/github_api"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "github-user-activity {user}",
		Example: "github-user-activity neel004",
		Short: "Github user activity is a cli tool used to grab the user's github activity.",
		RunE: FetchUserActivity,
		DisableFlagsInUseLine: true,
	}

	return cmd
}

func FetchUserActivity(cmd *cobra.Command, args []string) error {
	if len(args) == 0{
		return fmt.Errorf("expected user name.")
	}
	return api.GetUserActivity(args[0])
}
