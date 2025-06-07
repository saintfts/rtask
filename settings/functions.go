package settings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func parseLine(line string, line_n *int){
    defer func() {*line_n++}()
    if line == "" {
        return
    }
    switch line[0] { // skip comments
        case '#': // 1 option for now, switch-case for scalability
            return
    }
    param_and_val := strings.Split(line, "=")
    if len(param_and_val) == 0 {
        fmt.Printf("\ninvalid parameter in '.settings', line %d", *line_n)
        return
    }
    value := param_and_val[1]
    switch param_and_val[0] {
        case "PORT":
            AppSettings.Port = value
        case "DELIMITER":
            AppSettings.Delimiter = value
        case "IDEAS_PATH":
            AppSettings.TasksPath = value
        case "NEW_IDEAS_PATH":
            AppSettings.NewTasksPath = value
    }
}

func ParseSettings() { // Parses settings in ".settings" file and assignes to "settings" variable
    file, err := os.Open(".settings")
    defer file.Close()
    if err != nil {
        //defaultSettings()
        return
    }

    scanner := bufio.NewScanner(file)
    line_n := 1
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        parseLine(line, &line_n)
    }
}


