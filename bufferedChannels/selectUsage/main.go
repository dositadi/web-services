package main

import (
	"fmt"
	"time"
)

// Select allows us to lookup different channels more like using a switch statement
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 700)
		channel1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(time.Microsecond * 500)
		channel2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-channel1:
		fmt.Println(msg1)
	case msg2 := <-channel2:
		fmt.Println(msg2)
	}

	fmt.Println("Main exited!")
}
