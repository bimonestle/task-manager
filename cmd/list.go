package cmd

import (
	"fmt"
	"os"

	"github.com/bimonestle/go-exercise-projects/07.CLI-Task-Manager/task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Whty not take a vacation? 🏖")
			return
		}
		// fmt.Println("You have the following tasks: ", tasks) // TESTING
		for i, task := range tasks {
			fmt.Printf("%d. %s, Key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
