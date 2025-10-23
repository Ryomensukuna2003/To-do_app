/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "This will delete tasks",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid index.")
			return
		}
		todos := loadTodos()
		todos = append(todos[:index-1], todos[index:]...) // append concats 2 array so everything before index and everything after index
		saveTodos(todos);
		fmt.Println("Deleted Successfully");
		printTodos(todos)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
