package cmd

import (
	"fmt"
	"noting_webapp/parsers"
	"noting_webapp/server"
	"noting_webapp/settings"

	"github.com/spf13/cobra"
)



var startServer = &cobra.Command{
	Use:   "startServer",
    Short: "Add notes and view them on localhost:Port or *YourDeviceIP*:Port",
	Run: func(cmd *cobra.Command, args []string) {
        if port, err := cmd.Flags().GetString("port"); err == nil { // Checks if GetString is NOT succeeded
            settings.AppSettings.Port = port
        }

        ideas, err := parsers.GetTasksFromFile("tasks.txt") // Get tasks from file
        if err != nil {
            panic("ERROR: " + err.Error())
        }
        fmt.Println(len(ideas))
        server.StartServer(ideas, settings.AppSettings.Port)
    },
}
func init() {
	rootCmd.AddCommand(startServer)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
    startServer.Flags().StringP("port", "p", "3000", "Port for gofiber server")
}
