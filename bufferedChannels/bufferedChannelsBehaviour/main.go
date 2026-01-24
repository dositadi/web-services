package main

import (
	"fmt"
	"time"
)

// To understand this concept very well, See the channel as a temporary file
// The dataProducer() writes (sends) data into this file (channel) as long as it is not full
// The dataConsumer() reads data from this file as long as it is not empty

func dataProducer(ch chan int, id int) {
	for i := range 5 {
		ch <- i
		fmt.Printf("Producer %v sent: %v, channel length: %v\n", id, i, len(ch))
		time.Sleep(time.Millisecond * 50)
	}
}

func consumerFunction(ch chan int) {
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Printf("Consumer recieved: %v, channel length: %v\n", val, len(ch))
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	ch := make(chan int, 5)

	go dataProducer(ch, 1)
	go dataProducer(ch, 2)
	consumerFunction(ch)

	time.Sleep(time.Nanosecond)
	close(ch)

	for task := range ch {
		fmt.Printf("Main: Drained %v after close\n", task)
	}

	fmt.Println("Main: Exiting!")
}
