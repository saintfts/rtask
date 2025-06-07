package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"noting_webapp/settings"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)
var Mtx sync.Mutex
func ArrAtoi(strs []string) ([]int, error) {
    del_lines := make([]int, len(strs))
    for i, s := range strs {
        n, err := strconv.Atoi(s)
        if err != nil {
            return nil, err
        }
        del_lines[i] = n
    }
    return del_lines, nil
}
func RemoveTasksFromFile(path string, tasks_numbers ...int) error {
    tasks, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("couldn't open '%s': %v", path, err)
    }

    task_n := 1
    var lines []string
    scanner := bufio.NewScanner(tasks)
    for scanner.Scan() {
        line := scanner.Text()
        if slices.Contains(tasks_numbers, task_n){
            if line == settings.AppSettings.Delimiter {
                task_n++
            }
            continue
        } else {
            if line == settings.AppSettings.Delimiter {
                task_n++
            }
            lines = append(lines, line)
        }
    }
    tasks.Close()

    tasks_upd_path := path+".buff"
    tasks_upd, err := os.Create(tasks_upd_path)
    if err != nil {
        return fmt.Errorf("couldn't create '%s': %v", tasks_upd_path, err)
    }
    writer := bufio.NewWriter(tasks_upd)
    for _, line := range lines {
        _, err := writer.WriteString(line+"\n")
        if err != nil {
            return fmt.Errorf("error writing to file: %v", err)
        }
    }
    fmt.Println(lines)
    writer.Flush()
    tasks_upd.Close()

    remove_err := os.Remove(path)
    rename_err := os.Rename(tasks_upd_path, path)
    if remove_err != nil || rename_err != nil {
        if remove_err != nil {
            log.Println(remove_err)
        }
        if rename_err != nil {
            log.Println(rename_err)
        }
        return fmt.Errorf("couldn't swap %s with '%s'", tasks_upd_path, path)
    }
    return nil
}

func AppendTasksToFile(t, path string) {
    Mtx.Lock()
    defer Mtx.Unlock()
    tasks, err := os.OpenFile(path, os.O_APPEND, 0664)
    if err != nil {
        if os.IsNotExist(err) {
            log.Printf("file '%s' doesn't exist. Creating new...", path)
            tasks, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0664)
            if err != nil {
                log.Fatalf("couldn't create '%s': %v", path, err)
            }
        }else {
            log.Fatalln(err)
        }
    }
    defer tasks.Close()
    for strings.Contains(t, "\n\n") {
        t = strings.ReplaceAll(t, "\n\n", "\n")
    }
    if t[len(t)-1] == '\n'{
        t = t[:len(t)-1]
    }
    if _, err := tasks.WriteString(t+"\n```\n"); err != nil {
        log.Printf("writing err: %v", err)
    }
}
func AppendTasksFromTo(from, to string) error {
        newtasks, err := os.OpenFile(from, os.O_RDWR, 0664)
        if err != nil {
            if os.IsNotExist(err) {
                return fmt.Errorf("file %s doesn't exist: %v", from, err)
            }
            return fmt.Errorf("%v",err)
        }
        defer newtasks.Close()

        tasks, err := os.OpenFile(to, os.O_APPEND|os.O_WRONLY, 0664)
        if err != nil {
            if os.IsNotExist(err) {
                return fmt.Errorf("file %s doesn't exist: %v", to, err)
            }
            return fmt.Errorf("%v",err)
        }
        defer tasks.Close()

        if _, err = io.Copy(tasks, newtasks); err != nil {
            return fmt.Errorf("copy: %v", err)
        }

        if err = newtasks.Truncate(0); err != nil {
            return fmt.Errorf("truncate: %v", err)
        }
    return nil
}
