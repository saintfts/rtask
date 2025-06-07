package cmd

import (
	"fmt"
	"log"
	"noting_webapp/parsers"
	"noting_webapp/settings"

	"github.com/spf13/cobra"
)

var viewTasksCmd = &cobra.Command{
	Use:   "viewTasks",
	Short: "View tasks from tasks file",
	Run: func(cmd *cobra.Command, args []string) {
        tasks, err := parsers.GetTasksFromFile(settings.AppSettings.TasksPath)
        if err != nil {
            log.Fatalln(err)
        }
        for _, task := range tasks {
            fmt.Printf("%d) %s\n", task.Id, task.Text)
        }
	},
}

func init() {
	rootCmd.AddCommand(viewTasksCmd)

}
