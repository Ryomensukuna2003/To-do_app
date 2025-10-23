/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

// Main add cmd
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Used to add task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := args[0]
		todos := loadTodos()
		todos = append(todos, Todo{Task: tasks, Done: false})
		saveTodos(todos)
		fmt.Println("Added task: ", tasks)
	},
}

// This function is used to load todos in memory to apppend tasks
func loadTodos() []Todo {
	data, err := os.ReadFile("Todos.json")
	if err != nil {
		return []Todo{} // return empty todo list
	}

	var todos []Todo
	_ = json.Unmarshal(data, &todos) // what data to extract and where to put em
	return todos
}

// This function will format the todo data and add it to Todos.json
func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile("Todos.json", data, 0644)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
