package cmd

import (
	"../db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task on your TODO list as complete",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			err = fmt.Errorf("please provide a valid task ID. Error: %s", err)
			fmt.Println(err)
			os.Exit(1)
		}

		err = db.MarkTaskDone(id)
		if err != nil {
			err = fmt.Errorf("%d is an invalid ID", id)
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("You have completed task #%s", args[0])
	},
}