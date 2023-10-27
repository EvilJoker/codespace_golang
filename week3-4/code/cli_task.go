package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 枚举
type TaskStatus int

const (
	TASK_STATUS_UNDONE TaskStatus = iota
	TASK_STATUS_DONE
)

var TaskStatusNames = map[int]string{
	int(TASK_STATUS_UNDONE): "undone",
	int(TASK_STATUS_DONE):   "done",
}

// 结构体
type Task struct {
	id       uint32
	title    string
	desc     string
	deadline time.Time
	status   TaskStatus
}

type TaskManager struct {
	tasks []Task
}

type Storage interface {
	Load()

	AddTask(taskArr *[]Task, task *Task)
	GetTask(id uint32) *Task
	DelTask(id uint32)
	GetAllTasks() []Task
}

func (taskManager *TaskManager) AddTask(title string, desc string, deadline time.Time) *Task {
	task := &Task{
		id:       rand.Uint32(),
		title:    title,
		desc:     desc,
		deadline: deadline,
		status:   TASK_STATUS_UNDONE,
	}
	taskManager.tasks = append(taskManager.tasks, *task)
	fmt.Printf("Add Task: %d %s\n", task.id, task.title)
	return task
}

func (taskManager *TaskManager) GetTask(id uint32) *Task {
	for _, task := range taskManager.tasks {
		if task.id == id {
			fmt.Printf("Get Task: %d\n", id)
			return &task
		}
	}
	fmt.Printf("Get Task faiiled: %d\n", id)
	return nil
}

func (taskManager *TaskManager) DelTask(id uint32) {
	removeIndex := -1
	tasks := taskManager.tasks
	for i, task := range tasks {
		if task.id == id {
			removeIndex = i
			break
		}
	}
	if removeIndex >= 0 {
		tasks = append(tasks[:removeIndex], tasks[removeIndex+1:]...)
		fmt.Printf("Del Task: %d\n", id)
	} else {
		fmt.Printf("Task %d not found\n", id)
	}
}

func (taskManager *TaskManager) GetAllTasks() []Task {
	for _, task := range taskManager.tasks {
		fmt.Printf("Query Task: [%d, %s, %s, %s, %s]\n", task.id, task.title, task.desc, task.deadline.Format("2006-01-02 15:04:05"), TaskStatusNames[int(task.status)])
	}
	return taskManager.tasks
}
