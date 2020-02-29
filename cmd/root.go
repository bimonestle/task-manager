package cmd

import "github.com/spf13/cobra"

// Root command of Cobra lib
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
