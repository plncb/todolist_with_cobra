package cmd

import (
	"strconv"
	"todolist/internal/database"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a taks base on its id",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		err = database.CompleteTask(database.DB, id)
		return err
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
