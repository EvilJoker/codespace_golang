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

// 枚举
type TaskStatus int

const (
	TASK_STATUS_UNDONE TaskStatus = iota
	TASK_STATUS_DONE
)

// 结构体
type Task struct {
	title    string
	deadline time.Time
	desc     string
}

func AddTask(taskArr *[]Task, task *Task) {
	*taskArr = append(*taskArr, *task)
	fmt.Println("Add Task: ", task.title)

}

func QueryTask(taskArr *[]Task) {
	for _, task := range *taskArr {
		fmt.Println("Query Task: ", task.title)
	}
}

func main() {
	fmt.Println("Hello, CliTask!")
	var tasks []Task
	AddTask(&tasks, &Task{"first task", time.Now(), "1"})
	QueryTask(&tasks)

}
