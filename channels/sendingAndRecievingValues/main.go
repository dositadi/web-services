package main

import "fmt"

func Producer(ch chan int) {
	for i := range 5 {
		ch <- i
		fmt.Println("Produced the value: ", i)
	}
	// Close the channel!
	close(ch)
}

func Consumer(ch chan int) {
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel has been closed")
			break
		}
		fmt.Println("Consumed value: ", value)
	}
}

func main() {
	channel := make(chan int)

	go Producer(channel)

	Consumer(channel)
}
