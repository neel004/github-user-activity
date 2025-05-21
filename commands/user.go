package commands

import (
	"fmt"
	api "github.com/neel004/github-user-activity/github_api"
)

func UserActivity(user string, client *api.Client) error{
	if len(user) <= 0{
		return fmt.Errorf("user can not be empty.")
	}
	
	err := client.GetUserActivity(user)
	if err != nil{
		fmt.Println(err)
	}
	return nil
}
