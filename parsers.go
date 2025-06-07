package main
//
//import (
//	"bufio"
//	"fmt"
//	"noting_webapp/settings"
//	"os"
//	"strings"
//)
//
//func ParseTasks() ([]Idea, error) {
//    file, err := os.Open("tasks.txt")
//    defer file.Close()
//
//    if err != nil {
//        if os.IsNotExist(err) {
//            file, err = os.Create("tasks.txt")
//            if err != nil {
//                return []Idea{}, fmt.Errorf("Failed to create 'tasks.txt' %w", err)
//            }
//            return []Idea{}, nil
//        } else {
//            return []Idea{}, fmt.Errorf("Failed to open 'tasks.txt' %w", err)
//        }
//    }
//    var ideas []Idea
//
//    scanner := bufio.NewScanner(file)
//    scanner.Scan()
//    var idea Idea = scanner.Text()
//    for scanner.Scan() {
//        line := scanner.Text()
//        if strings.TrimSpace(line) == settings.AppSettings.Delimiter {
//            ideas = append(ideas, idea)
//            if scanner.Scan(){
//                idea = scanner.Text()
//            }
//        } else {
//            idea = idea+"\n"+line
//        }
//    }
//    ideas = append(ideas, idea)
//
//    return ideas, nil
//    //todo...
//}
