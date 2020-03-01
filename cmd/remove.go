package cmd

import (
	"fmt"
	"strconv"

	"github.com/bimonestle/go-exercise-projects/07.CLI-Task-Manager/task/db"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete the task",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Remove called")
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		// 1.	Get all tasks
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}
		// 2.	Iterate over the ids
		// 	2a. condition: id should be in the list of ids
		// 	2b. condition: id less than equal to zero and greater than the length of tasks is invalid
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			// 3.	tasks[id -1]
			task := tasks[id-1]

			// 4.	DeleteTask(key)
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to remove \"%d\": \"%s\". Error: %s\n", id, task.Value, err)
			} else {
				fmt.Printf("Successfully deleted \"%d\": \"%s\"!", id, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
