package server

import (
	"encoding/json"
	"fmt"
	"log"
	"noting_webapp/settings"
	. "noting_webapp/structs"
	"os"
	"reflect"
	"strings"
	_"os/signal"
	_"context"
	_"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer(tasks []Task, port string){
        if _, err := os.Stat(settings.AppSettings.NewTasksPath); os.IsNotExist(err) {
            log.Printf("new tasks file '%s' doesn't exist. Creating new...", settings.AppSettings.NewTasksPath)
            if _, err := os.Create(settings.AppSettings.NewTasksPath); err != nil {
                log.Fatalf("couldn't create '%s'. Aborting...", err)
            }
        }




    /// In case if you want to make the new tasks join new_tasks file instead of tasks.txt
    //ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
        //go func(){
        //    <-(ctx).Done()
        //    stop()
        //    AppendTasksFromTo(settings.AppSettings.NewTasksPath, settings.AppSettings.TasksPath)
        //    fmt.Println("ctrl c called")
        //    os.Exit(1)
        //}()


        html_engine := html.New("./views", ".html")
        html_engine.AddFunc("add", func(a, b int) int {
            return a+b
        })
        app := fiber.New(fiber.Config{
                Views: html_engine,
            })

        app.Static("/", "./public")
        app.Get("/", func(c *fiber.Ctx) error {
            return c.Render("index", fiber.Map{
                "Tasks": tasks,
            })
        })

        app.Post("/api/send_task", func(c *fiber.Ctx) error {
            task := struct {
                Text string `json:"text"`
            }{}

            if err := c.BodyParser(&task); err != nil || len(task.Text) == 0 {
                return c.SendStatus(fiber.StatusBadRequest)
            }
            AppendTasksToFile(task.Text, settings.AppSettings.TasksPath)
            Tasks_n++
            tasks = append(tasks, Task{
                Id: Tasks_n,
                Text: strings.ReplaceAll(task.Text, "\n", ". "),
            })


            return c.SendStatus(fiber.StatusAccepted)
        })


        app.Get("/api/add_task", func(c *fiber.Ctx) error {
            Tasks_n++
            tasks = append(tasks, Task{Id: Tasks_n, Text: "NewIdea"})
            return c.SendStatus(fiber.StatusAccepted)
        })


        app.Post("/api/remove_task", func(c *fiber.Ctx) error {
            task_n := struct {
                N int `json:"n"`
            }{}
            if err := c.BodyParser(&task_n); err != nil {
                log.Printf("remove lines failed3:")
                return c.Status(fiber.StatusAlreadyReported).JSON(fiber.Map{"Error": "Wrong object's keys"})
            }
            log.Println(task_n)
            if task_n.N < 1 || task_n.N > len(tasks) {
                log.Printf("remove lines failed2: %v", task_n.N)
                return c.SendStatus(fiber.StatusBadRequest)
            }
            if err := RemoveTasksFromFile(settings.AppSettings.TasksPath, task_n.N); err != nil {
                log.Printf("remove lines failed: %v", err)
            }
            tasks = append(tasks[:task_n.N-1], tasks[task_n.N:]...)
            return c.SendStatus(fiber.StatusAccepted)
        })


        app.Post("/api/get_new_tasks", func(c *fiber.Ctx) error {
            body := c.Body()
            var data map[string]interface{}
            if err:= json.Unmarshal(body, &data); err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error:": "Wrong JSON"})
            }
            log.Println(data)
            if len(data) == 0 {
                return c.Status(fiber.StatusCreated).JSON(tasks)
            }
            last_id := int(data["id"].(float64))
            log.Println( last_id, reflect.TypeOf(last_id))

            if last_id > int(len(tasks)) {
                log.Printf("\nUnusual last_id (%d vs %d) sent from IP %s", last_id, len(tasks), c.IP())
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
            } else if last_id < int(len(tasks)) {
                return c.Status(fiber.StatusCreated).JSON(tasks[last_id:]) 
            }

            return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
        })


        app.Get("/api/tasks", func(c *fiber.Ctx) error {
            fmt.Println(tasks)
            return c.JSON(tasks)
        })

        app.Listen(":" + settings.AppSettings.Port)
}



