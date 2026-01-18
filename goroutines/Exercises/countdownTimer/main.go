package main

import (
	"fmt"
	"time"
)

func CountDown(id, start int) {
	for i := range start {
		count := i
		time.Sleep(1 * time.Second)
		fmt.Println("Timer ", id, ": ", count)
	}
}

func main() {
	fmt.Println("Main routine started!")

	go CountDown(1, 4)
	go CountDown(2, 6)

	time.Sleep(10 * time.Second)
	fmt.Println("Main routine ended!")
}
