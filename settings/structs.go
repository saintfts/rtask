package settings

type Settings struct {
    Port string
    TasksPath string
    NewTasksPath string
    Delimiter string
}

var AppSettings Settings = Settings {
    Port: "3000",
    TasksPath: "tasks.txt",
    NewTasksPath: "new_tasks.txt",
    Delimiter: "```",
}
