/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a readme file",
	Long: `create a readme file. For example:
		readme create -t "My Project" -d "This is my project" -a "John Doe" -l "MIT" -u "`,
	Run: func(cmd *cobra.Command, args []string) {
		readmeTitle, _ := cmd.Flags().GetString("title")
		readmeDescription, _ := cmd.Flags().GetString("description")
		fmt.Println("Title:", readmeTitle)
		fmt.Println("Description:", readmeDescription)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "Title of the project")
	createCmd.Flags().StringP("description", "d", "", "Description of the project")
}

