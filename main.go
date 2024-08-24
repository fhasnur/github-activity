package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "github-activity [username]",
		Short: "CLI to display user GitHub activity",
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			fmt.Println(username)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
