package main

import "fmt"

func main() {
	bufferedChannel := make(chan int, 3)

	fmt.Println("Channels capacity: ", cap(bufferedChannel))

	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3

	firstValue := <-bufferedChannel
	fmt.Println(firstValue)

	bufferedChannel <- 4

	fmt.Println("value: ", <-bufferedChannel)
	fmt.Println("value: ", <-bufferedChannel)
	fmt.Println("value: ", <-bufferedChannel)
	fmt.Println(len(bufferedChannel))
}
