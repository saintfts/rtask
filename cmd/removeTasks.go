/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"noting_webapp/server"
	"noting_webapp/settings"

	"github.com/spf13/cobra"
)
var removeTasksCmd = &cobra.Command{
	Use:   "removeTasks",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
        del_lines, err := server.ArrAtoi(args)
        if err != nil {
            log.Fatalf("bad arguments sequence: %v", err)
        }
        server.RemoveTasksFromFile(settings.AppSettings.TasksPath, del_lines...)

	},
}

func init() {
	rootCmd.AddCommand(removeTasksCmd)

}
