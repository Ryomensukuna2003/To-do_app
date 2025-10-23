/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks",

	Run: func(cmd *cobra.Command, args []string) {
		todos := loadTodos()
		printTodos(todos)
	},
}

func printTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No task found!!!")
		return
	}

	idStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#C792EA")).Width(4)
	taskStyle := lipgloss.NewStyle().Padding(0, 1)
	doneStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#43D787")).Bold(true)
	notDoneStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6B6B"))

	var lines []string
	for i, t := range todos {
		status := notDoneStyle.Render("Not Done")
		if t.Done {
			status = doneStyle.Render("Done")
		}
		line := fmt.Sprintf("%s %s %s",
			idStyle.Render(fmt.Sprintf("%d.", i+1)),
			taskStyle.Render(t.Task),
			status,
		)
		lines = append(lines, line)
	}

	fmt.Println(lipgloss.JoinVertical(lipgloss.Top, lines...))
}


func init() {

	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
