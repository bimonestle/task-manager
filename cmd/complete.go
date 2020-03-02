package cmd

import (
	"fmt"
	"os"

	"github.com/bimonestle/go-exercise-projects/07.CLI-Task-Manager/task/db"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Lists all the completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
		tasks, err := db.AllCompleted()
		if err != nil {
			fmt.Println("Someting went wrong: ", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You don't have any completed tasks. Be productive!")
			return
		}
		for i, task := range tasks {
			fmt.Printf("You have completed: %d. %s, Key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(completeCmd)
}
