package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"todolist/internal/database"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := database.GetAllTasks(database.DB)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)
		fmt.Fprintln(w, "Task ID\tDescription\tCreated\tDue By\tCompleted")
		for _, task := range tasks {
			fmt.Fprintf(
				w,
				"%d\t%s\t%s\t%s\t%t\n",
				task.ID,
				task.Description,
				timediff.TimeDiff(task.CreatedAt),
				timediff.TimeDiff(task.DueBy),
				task.IsCompleted,
			)
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
