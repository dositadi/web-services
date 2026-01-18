package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main routine started!")

	for i := range 3 {
		// Capture the value of i
		count := i

		go func() {
			time.Sleep(time.Duration(count) * 100 * time.Millisecond)
			fmt.Println("Anonymous goroutine, ", count)
		}()
	}

	fmt.Println("Main finished launching go routine!")

	time.Sleep(1 * time.Second)
	fmt.Println("Main routine ended!")
}
