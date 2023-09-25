/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "readme",
	Short: "readme is a CLI for generating README.md files for your projects",
	Long: `readme is a CLI for generating README.md files for your projects. For example: 
		readme init -t "My Project" -d "This is my project" -a "John Doe" -l "MIT" -u "https://github.com/johndoe/myproject" -e "`,
	
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


