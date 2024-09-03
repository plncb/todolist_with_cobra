package cmd

import (
	"log"
	"time"
	"todolist/internal/database"

	"github.com/spf13/cobra"
)

var dueBy string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		const dateFormat = "2006-01-02" // Format: YYYY-MM-DD
		parsedDate, err := time.Parse(dateFormat, dueBy)
		if err != nil {
			log.Fatalf("Invalid date format. Please use YYYY-MM-DD.")
			return err
		}

		_, err = database.CreateTask(database.DB, args[0], parsedDate)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&dueBy, "date", "d", "9999-12-31", "Date in YYYY-MM-DD format")
	addCmd.MarkFlagRequired("date")
	rootCmd.AddCommand(addCmd)
}
