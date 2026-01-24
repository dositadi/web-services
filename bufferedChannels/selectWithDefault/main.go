package main

import (
	"fmt"
	"time"
)

// Select with default statements never block

func main() {
	channel1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Millisecond * 500)
		channel1 <- "data available!"
	}()

	select {
	case msg1 := <-channel1:
		fmt.Println("Recieved Message: ", msg1)
	default:
		fmt.Println("No data is available!")
	}

	time.Sleep(time.Millisecond * 500)

	select {
	case msg1 := <-channel1:
		fmt.Println("Recieved Message: ", msg1)
	default:
		fmt.Println("No data is available!")
	}
}
