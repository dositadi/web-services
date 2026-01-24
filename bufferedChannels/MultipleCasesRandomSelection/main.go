package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan string)
	chB := make(chan string)
	var chC chan string
	fmt.Println(chC)

	chAInput := []string{"Good", "Better", "Best"}
	chBInput := []string{"Grow", "on", "fast!"}

	for i := 0; i < 3; i++ {
		go func() {
			time.Sleep(time.Millisecond * 100)
			chA <- chAInput[i]
		}()

		go func() {
			time.Sleep(time.Millisecond * 100)
			chB <- chBInput[i]
		}()
	}

	select {
	case msg1 := <-chA:
		fmt.Println("From chA: ", msg1)
	case msg2 := <-chB:
		fmt.Println("From chB: ", msg2)
	case <-time.After(time.Millisecond * 700):
		fmt.Println("Timed out!")
	}
}
