package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID    int
	Value int
}

type Result struct {
	ID    int
	Value int
}

func ProcessTask(current int, tasks <-chan Task, results chan<- Result) {
	for task := range tasks {
		fmt.Println("Processing task: ", task.ID, "Value: ", task.Value)
		time.Sleep(time.Millisecond * 100)
		results <- Result{ID: task.ID, Value: task.Value * 2}
	}
	fmt.Println("Processed task ", current)
}

func main() {
	var numWorkers = 3
	var amountOfTasks = 10

	tasks := make(chan Task, amountOfTasks/2)
	results := make(chan Result, amountOfTasks)

	for i := range numWorkers {
		go ProcessTask(i, tasks, results)
	}

	for i := range amountOfTasks {
		tasks <- Task{ID: i, Value: i + 1}
	}
	close(tasks)

	for i := 0; i < amountOfTasks; i++ {
		result := <-results
		fmt.Println("Recieved task: ", result.ID, " with value: ", result.Value)
	}
	close(results)
}
