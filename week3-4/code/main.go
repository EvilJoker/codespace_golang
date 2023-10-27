package main

/*
初始化：
创建 go mod init example.com/cli_task
运行 go run .

*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func doCommand(taskManager *TaskManager, action string, body string) {
	switch action {
	case "add":
		// go run . --command add -body "{\"title\":\"one\", \"desc\":\"demo task\", \"deadline\":\"2006-01-02 15:04:05\"}"
		if len(body) == 0 {
			fmt.Println("miss body")
			break
		}

		var data map[string]interface{}
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			fmt.Println("json parse error", err)
			return
		}

		title := data["title"].(string)
		if title == "" {
			fmt.Println("miss title")
			return
		}

		desc := data["desc"].(string)
		if desc == "" {
			fmt.Println("miss desc")
			return
		}
		deadline := data["deadline"].(string)
		if deadline == "" {
			fmt.Println("miss deadline")
			return
		}
		deadline_t, err1 := time.Parse("2006-01-02 15:04:05", deadline)
		if err1 != nil {
			fmt.Println("deadline parse error", err1)
			return
		}

		taskManager.AddTask(title, desc, deadline_t)

	case "delete":
		if len(body) == 0 {
			fmt.Println("miss body")
			break
		}
		id, err := strconv.ParseUint(body, 10, 32)
		if err != nil {
			fmt.Println("id parse error", err)
		}
		taskManager.DelTask(uint32(id))

	case "list":
		taskManager.GetAllTasks()
	default:
		fmt.Println("[usage] cli_task -command [add|delete|list] -body [1|{}]")
	}
}

func main() {

	var taskManager TaskManager
	taskManager.Load()

	commandPtr := flag.String("command", "", "[usage] cli_task -command [add|delete|list] -body [1|{}]")
	bodyPtr := flag.String("body", "", "when add: \"{\"title\":\"one\", \"desc\":\"demo task\", \"deadline\":\"2006-01-02 15:04:05\"}\"\n"+
		"when delete: {'id':'1'}\n")
	flag.Parse()

	doCommand(&taskManager, *commandPtr, *bodyPtr)
	taskManager.Save()

}
