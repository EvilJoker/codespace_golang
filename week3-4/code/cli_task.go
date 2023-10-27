package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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
	Id       uint32     `json:"Id"`
	Title    string     `json:"Title"`
	Desc     string     `json:"Desc"`
	Deadline time.Time  `json:"Deadline"`
	Status   TaskStatus `json:"Status"`
}

type TaskManager struct {
	Tasks []Task `json:"tasks"`
}

type Storage interface {
	Load()
	Save()
	AddTask(taskArr *[]Task, task *Task)
	GetTask(id uint32) *Task
	DelTask(id uint32)
	GetAllTasks() []Task
}

func (taskManger *TaskManager) Load() {
	file, err := os.OpenFile("tasks.json", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&taskManger)
	if err != nil {
		panic(err)
	}
}

func (taskManager *TaskManager) Save() {
	jsonStr, err0 := json.Marshal(taskManager)

	if err0 != nil {
		panic(err0)
	}
	file, err1 := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err1 != nil {
		panic(err1)
	}
	defer file.Close()
	_, err2 := file.WriteString(string(jsonStr))
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("save in tasks.json success")
}

func (taskManager *TaskManager) AddTask(title string, desc string, deadline time.Time) *Task {
	task := &Task{
		Id:       rand.Uint32(),
		Title:    title,
		Desc:     desc,
		Deadline: deadline,
		Status:   TASK_STATUS_UNDONE,
	}
	taskManager.Tasks = append(taskManager.Tasks, *task)
	fmt.Printf("Add Task: %d %s\n", task.Id, task.Title)
	return task
}

func (taskManager *TaskManager) GetTask(id uint32) *Task {
	for _, task := range taskManager.Tasks {
		if task.Id == id {
			fmt.Printf("Get Task: %d\n", id)
			return &task
		}
	}
	fmt.Printf("Get Task faiiled: %d\n", id)
	return nil
}

func (taskManager *TaskManager) DelTask(id uint32) {
	removeIndex := -1
	tasks := taskManager.Tasks
	for i, task := range tasks {
		if task.Id == id {
			removeIndex = i
			break
		}
	}
	if removeIndex >= 0 {
		taskManager.Tasks = append(tasks[:removeIndex], tasks[removeIndex+1:]...)
		fmt.Printf("Del Task: %d\n", id)
	} else {
		fmt.Printf("Task %d not found\n", id)
	}
}

func (taskManager *TaskManager) GetAllTasks() []Task {
	for _, task := range taskManager.Tasks {
		fmt.Printf("Query Task: [%d, %s, %s, %s, %s]\n", task.Id, task.Title, task.Desc, task.Deadline.Format("2006-01-02 15:04:05"), TaskStatusNames[int(task.Status)])
	}
	return taskManager.Tasks
}
