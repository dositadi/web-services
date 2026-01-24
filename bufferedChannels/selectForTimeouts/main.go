package main

import (
	"fmt"
	"time"
)

func Worker(done chan bool) {
	fmt.Println("Doing some work!")
	time.Sleep(1 * time.Second)
	fmt.Println("Work completed!")
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go Worker(done)

	select {
	case <-done:
		fmt.Println("Main: Task has been completed!")
	case <-time.After(2 * time.Second):
		fmt.Println("Main: Task timed out")
	}
	fmt.Println("Main exiting!")
}
