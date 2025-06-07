/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "fmt"
	"noting_webapp/cmd"
	_ "noting_webapp/parsers"
	_ "noting_webapp/server"
	"noting_webapp/settings"
)


func main() {
    settings.ParseSettings()

    //a, _ := parsers.GetTasksFromFile("tasks.txt")
    //server.RemoveTasksFromFile(settings.AppSettings.TasksPath, 1,3)
    cmd.Execute()
}
