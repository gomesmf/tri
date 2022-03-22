/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/gomesmf/tri/todo"
	"github.com/spf13/cobra"
)

func addRun(cmd *cobra.Command, args []string) {
	items := []todo.Item{}

	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}

	err := todo.SaveItems("/Users/mfgomes/.tridos.json", items)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
