package cmd

import (
	"../db"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks.")
		} else {
			fmt.Println("You have the following tasks:")
		}
		for _, task := range tasks {
			var completed string
			if task.Done {
				completed = "☑"
			} else {
				completed = "☐"
			}

			fmt.Printf("%s %d: %s\n", completed, task.ID, task.Description)
		}
	},
}