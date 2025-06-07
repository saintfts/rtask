package parsers

import (
	"bufio"
	"fmt"
	"noting_webapp/server"
	"noting_webapp/settings"
	. "noting_webapp/structs"
	"os"
	"strings"
)

func ParseTasksFromFile(file *os.File) []Task {
    var tasks []Task
    scanner := bufio.NewScanner(file)
    var task strings.Builder

    for scanner.Scan() {
        line := scanner.Text()
        if trimmed_line := strings.TrimSpace(line); trimmed_line == settings.AppSettings.Delimiter {
            tasks = append(tasks, Task {Id: server.Tasks_n, Text: task.String()})
            task.Reset()
            server.Tasks_n++;
        } else if trimmed_line=="" {
            continue
        } else {
            if task.Len()==0 { // if new line
                task.WriteString(line)
            } else {
                task.WriteString(". "+line)
            }
        }
    }
    if task.Len() != 0 {
        tasks = append(tasks, Task {Id: server.Tasks_n, Text: task.String()})
    }
    return tasks
}

func GetTasksFromFile(path string) ([]Task, error) {
    file, err := os.Open(path)

    if err != nil {
        if os.IsNotExist(err) {
            file, err = os.Create(path)
            if err != nil {
                return []Task{}, fmt.Errorf("cannot to create %s %w", path, err)
            }
            return []Task{}, nil
        } else {
            return []Task{}, fmt.Errorf("cannot to open %s %w", path, err)
        }
    }
    defer file.Close()
    return ParseTasksFromFile(file), nil
}
