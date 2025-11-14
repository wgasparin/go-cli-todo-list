package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {

	action := os.Args[1]

	path := "./data/todo.json"

	tasks := readTasksFromDataSource(path)

	switch action {
	case "add":
		task := os.Args[2]

		tasks, id := addTask(task, tasks)

		fmt.Printf("Taks id %v has been added", id)

		writeTasksToDataSource(path, tasks)

		return
	case "list":
		fmt.Println("Listing all tasks")
		listTasks(tasks)
		return
	case "complete":
		taskIdArg := os.Args[2]
		taskId, _ := strconv.Atoi(taskIdArg)

		tasks = completeTask(taskId, tasks)

		writeTasksToDataSource(path, tasks)
		fmt.Println("Task has been completed")

		return
	case "remove":
		taskIdArg := os.Args[2]
		taskId, _ := strconv.Atoi(taskIdArg)

		tasks = removeTask(taskId, tasks)

		writeTasksToDataSource(path, tasks)

		fmt.Println("Task has been removed")

		return
	}

	fmt.Println("Unknown option:", action)
	fmt.Println("Options:")
	fmt.Println("\tlist")
	fmt.Println("\tadd taskdescription")
	fmt.Println("\tremove taskId")
	fmt.Println("\tcomplete taskId")
}

func readTasksFromDataSource(path string) (tasks []myTask) {
	file, _ := os.ReadFile(path)

	_ = json.Unmarshal(file, &tasks)

	return tasks
}

func writeTasksToDataSource(path string, tasks []myTask) {
	file, _ := json.Marshal(tasks)

	os.WriteFile(path, file, 0644)
}

func getMaxId(tasks []myTask) int {
	maxId := 0

	for _, task := range tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	return maxId
}

func addTask(task string, tasks []myTask) (newTasks []myTask, id int) {

	id = getMaxId(tasks) + 1

	newTask := myTask{
		ID:        id,
		Task:      task,
		Completed: false,
	}

	newTasks = append(tasks, newTask)

	return
}

func listTasks(tasks []myTask) {
	for _, todo := range tasks {

		done := ""

		if todo.Completed {
			done = "(done)"
		}

		fmt.Printf("[%v] Task: %v %v\n", todo.ID, done, todo.Task)
	}
}

func removeTask(id int, tasks []myTask) []myTask {

	for index, task := range tasks {
		if task.ID == id {
			return append(tasks[:index], tasks[index+1:]...)
		}
	}

	return tasks
}

func completeTask(id int, tasks []myTask) []myTask {

	for index, task := range tasks {
		if task.ID == id {
			tasks[index].Completed = true
			break
		}
	}

	return tasks
}

type myTask struct {
	ID        int
	Task      string
	Completed bool
}
