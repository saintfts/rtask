package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "noting_webapp",
    Short: "Add notes and view them on localhost:Port or *YourDeviceIP*:Port",
}


func Execute() {
    rootCmd.AddCommand(startServer)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("port", "p", "3000", "Port for the app")
}


