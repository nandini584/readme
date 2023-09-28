/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"io/ioutil"
	
	
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a readme file",
	Long: `create a readme file. For example:
		readme create -t "My Project" -d "This is my project" -a "John Doe" -l "MIT" -u "`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err:= os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("current working directory", currentDir)
		
		readmeTitle, _ := cmd.Flags().GetString("title")
		readmeDescription, _ := cmd.Flags().GetString("description")
		fmt.Println("Title:", readmeTitle)
		fmt.Println("Description:", readmeDescription)

		markdownContent := generateMarkdown (readmeTitle, readmeDescription)
		// data, err := ioutil.ReadFile("readme.md")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// markdownContent := string(data)
		// title, description := extract(markdownContent)
		// fmt.Println("Title:", title)
		// fmt.Println("Description:", description)

		
		// file, err:= os.Create("readme.md")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		err2:= ioutil.WriteFile("nandini.md", []byte(markdownContent),0644)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		// defer file.Close()
		fmt.Println("File created successfully")
	},

}

func generateMarkdown(title, description string) string {
	markdown:=fmt.Sprintf("# %s\n\n%s\n", title, description)
	return markdown
}
func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "Title of the project")
	createCmd.Flags().StringP("description", "d", "", "Description of the project")
}

