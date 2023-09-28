/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
"fmt"
"os"
"github.com/spf13/cobra"
"io/ioutil")
var rootCmd = &cobra.Command{Use: "readme-cli"}
func main() {
	dir,err:=os.Getwd()
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println("cwd", dir)
	file,err:= os.Create("readme.md")
		if err!=nil {
			fmt.Println(err)
			return
		}
	err2:= ioutil.WriteFile("readme.md",[]byte("hello world"),0644)
	if err2!= nil {
		fmt.Println(err2)
	}
	defer file.Close();
	if err:= rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}
