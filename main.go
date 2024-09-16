package main

import (
	"fmt"
	"os"

	"github.com/fhasnur/github-activity/pkg/activity"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "github-activity [username]",
		Short: "CLI to display user GitHub activity",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			err := activity.FetchGithubActivity(username)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
