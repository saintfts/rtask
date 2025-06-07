package cmd

import (
	"log"
	"noting_webapp/server"
	"noting_webapp/settings"

	"github.com/spf13/cobra"
)

var addNotesCmd = &cobra.Command{
	Use:   "addTasks",
	Short: "Add tasks from new_tasks file",
	Run: func(cmd *cobra.Command, args []string) {
        if err := server.AppendTasksFromTo(settings.AppSettings.NewTasksPath, settings.AppSettings.TasksPath); err != nil {
            log.Fatalln(err)
        }
	},
}
var addFromText = &cobra.Command{
	Use:   "text",
	Short: "Add note from text",
	Run: func(cmd *cobra.Command, args []string) {
        for _, task := range args {
            server.AppendTasksToFile(task, settings.AppSettings.TasksPath)
        }
    },
}
func init() {
	addNotesCmd.AddCommand(addFromText)
    rootCmd.AddCommand(addNotesCmd)

	addNotesCmd.Flags().BoolP("with_file", "", false, "Add new tasks from a file")

	addNotesCmd.Flags().StringP("file", "f", "new_tasks.txt", "Path for the new tasks file")
}
