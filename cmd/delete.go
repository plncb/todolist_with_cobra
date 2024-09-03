package cmd

import (
	"fmt"
	"strconv"
	"todolist/internal/database"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task base on its id",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		err = database.DeleteTask(database.DB, id)
		if err == nil {
			fmt.Printf("Task %d successfully deleted", id)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
