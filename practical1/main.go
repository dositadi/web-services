package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a cli programme that allows a user create tasks which are stored in a text.txt file which serves as the data base and as such the user can update this to mark the task he or she has successfully done or completed
*/

// NewTask struct for creating new tasks
type Task struct {
	Description string `json: "description"`
	Completed   bool   `json: "completed"`
}

// Constructor for NewTask
func CreateNewTask(desc string) Task {
	return Task{
		Description: desc,
		Completed:   false,
	}
}

// Display task to user in a string friendly
func (t Task) String() string {
	status := " "
	if t.Completed {
		status = "x"
	}
	return fmt.Sprintf("[%s] %s", status, t.Description)
}

// Function to add task to the data base
func AddTask(tasks []Task, description string) []Task {
	newTask := CreateNewTask(description)
	tasks = append(tasks, newTask)
	return tasks
}

// Function to list out all the tasks to the user
func ListTasks() {
	var tasksList []Task
	tasksList = UnmarshalJson(tasksList)
	if len(tasksList) == 0 {
		fmt.Println("You dont have any task yet")
	}

	for i, task := range tasksList {
		fmt.Println(i+1, ". "+task.String())
	}
}

// Function to update task if it has been completed successfully
func Complete(tasks *[]Task, index int) error {
	slice := *tasks

	if len(slice) < 1 {
		return fmt.Errorf("No task yet!")
	}

	if index < 0 || index >= len(slice) {
		return fmt.Errorf("Index(%v) out of bounds: 0 to %v", index+1, len(slice))
	}

	slice[index].Completed = true

	MarshalJsonToFile(slice)
	return nil
}

// Function to delete a task from the task database
func DeleteTask(tasks []Task, index int) error {
	if tasks == nil {
		return fmt.Errorf("No existing task yet!")
	}

	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("Index(%v) out of bounds: 0 to %v", index+1, len(tasks))
	}

	var slice []Task

	for i, task := range tasks {
		if i != index {
			slice = append(slice, task)
		}
	}

	MarshalJsonToFile(slice)
	return nil
}

func OpenAndReadFile(filePath string) ([]byte, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Read file error: %w", err)
	}

	fileInfo, err1 := file.Stat()
	if err1 != nil {
		return nil, fmt.Errorf("Read file error: %w", err1)
	}

	fileBuffer := make([]byte, fileInfo.Size())

	_, err2 := file.Read(fileBuffer)
	if err2 != nil {
		return nil, fmt.Errorf("Read file error: %w", err2)
	}

	return fileBuffer, nil
}

func OpenAndWriteToFile(filePath string, tasks []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Write to file error: %w", err)
	}

	_, err1 := file.WriteString(string(tasks))
	if err1 != nil {
		return fmt.Errorf("Write to file error: %w", err1)
	}
	return nil
}

func UnmarshalJson(tasks []Task) []Task {
	taskStore, err := OpenAndReadFile(FILEPATH)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err1 := json.Unmarshal(taskStore, &tasks)
	if err1 != nil {
		fmt.Println("Unable to unmarshal the json data")
	}
	return tasks
}

func MarshalJsonToFile(tasks []Task) {
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Unable to marshall tasks to JSON!: ", err)
		return
	}

	err1 := OpenAndWriteToFile(FILEPATH, tasksJSON)
	if err1 != nil {
		fmt.Println(err)
	}
}

const FILEPATH = "practical1/task.txt"

func main() {
	var tasks []Task
	tasks = UnmarshalJson(tasks)

	if len(os.Args) < 2 {
		printUsage()
		return
	}
	args := os.Args[1:]
	var command string

	command = args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Add Task: go run main.go add <description>")
			os.Exit(1)
		}

		description := strings.Join(args[1:], " ")

		if description == "" {
			fmt.Println("Error: Description cannot be empty!")
			fmt.Println()
			fmt.Println("Add Task: go run main.go add <description>")
			os.Exit(1)
		}

		tasks = AddTask(tasks, description)
		MarshalJsonToFile(tasks)
		fmt.Println("Successfully added ", description, " to tasks")

	case "list":
		ListTasks()

	case "complete":
		if len(args) < 2 {
			fmt.Println("Complete Task: go run main.go complete <task position>")
			os.Exit(1)
		}

		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid position value.\nHint: Position should be a number")
			os.Exit(1)
		}
		err1 := Complete(&tasks, index-1)
		if err1 != nil {
			fmt.Println("Error: ", err1)
			os.Exit(1)
		} else {
			fmt.Println("Congratulations on completing `", tasks[index-1], "` task. Your progress has been recorded successfully!!!")
		}
	case "delete":
		if len(args) < 2 {
			fmt.Println("Delete Task: go run main.go delete <task position>")
			os.Exit(1)
		}

		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid position value.\nHint: Position should be a number")
			os.Exit(1)
		}

		err2 := DeleteTask(tasks, index-1)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			os.Exit(1)
		} else {
			fmt.Println("You have deleted task `", tasks[index-1], "` successfully!")
		}
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("---System Usage---")
	fmt.Println("Use the following Commands:")
	fmt.Println("Add Task: go run main.go add <description>")
	fmt.Println("List Tasks: go run main.go list")
	fmt.Println("Complete Task: go run main.go complete <task position>")
	fmt.Println("Delete Task: go run main.go delete <task position>")
}
