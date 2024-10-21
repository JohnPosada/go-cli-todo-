package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	task "github.com/JohnPosada/go-cli-todo/tasks"
	"github.com/google/uuid"
)

func main() {

	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Print("error :", err)
		return

	}
	defer file.Close()

	tasks := task.LoadTask(file)

	switch os.Args[1] {
	case "list":
		task.List(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter task name:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.Add(tasks, name)
		task.Save(file, tasks)
		task.List(tasks)

	case "done":
		if len(os.Args) != 3 {
			fmt.Println("invalid command, enter id")
			return
		}
		id := os.Args[2]
		uuid, err := uuid.Parse(id)
		if err != nil {
			panic(err)
		}

		tasks = task.Done(tasks, uuid)
		task.Save(file, tasks)
		task.List(tasks)
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("invalid command, enter id")
			return
		}
		id := os.Args[2]
		uuid, err := uuid.Parse(id)
		if err != nil {
			panic(err)
		}

		tasks = task.Delete(tasks, uuid)
		task.Save(file, tasks)
		task.List(tasks)
	default:
		fmt.Println("invalid command, enter a valid command")
		fmt.Println("list-- list all tasks")
		fmt.Println("add-- add new task")
		fmt.Println("done-- mark task as done")
		fmt.Println("delete-- delete task")

	}

}
