package main

import (
	"noting_webapp/parsers"
	_ "noting_webapp/structs"
	"testing"
)

var test_tasks []string = []string{
    "Test1",
    `Test1
b`,
}

func TestParseTasks(t *testing.T){
    tasks, err := parsers.GetTasksFromFile("tests/test_tasks.txt")

    if err != nil {
        t.Fail()
    }
    t.Log(tasks, len(tasks))
    for i := 0; i < max(len(tasks), len(test_tasks)); i++ {
        if tasks[i].Text != test_tasks[i] {
            t.Error("test_tasks is not equal to tasks in test_tasks.txt")
        }
    }
}
