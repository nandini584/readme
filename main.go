/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import ("github.com/nandini584/readme-cli/cmd"
"fmt"
"os"
"github.com/spf13/cobra")
var rootCmd = &cobra.Command{Use: "readme-cli"}
func main() {
	if err:= rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
