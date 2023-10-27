package main

/*
初始化：
创建 go mod init example.com/cli_task
运行 go run .

*/

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, CliTask!")
	var taskManager TaskManager

	task := taskManager.AddTask("one", "first task", time.Now())
	taskManager.GetAllTasks()
	taskManager.DelTask(task.id)
	taskManager.GetAllTasks()

}
