package main

import (
	"fmt"
	"time"
)

func DoWork(values []int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range values {
			input := values[i]
			out <- input
		}
		close(out)
	}()
	return out
}

func RecieveWork(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func main() {
	slice := []int{1, 2, 3, 5, 6}
	dataToChannel := DoWork(slice)
	finalChannel := RecieveWork(dataToChannel)

	var result []int
	for i := range finalChannel {
		result = append(result, i)
	}
	fmt.Println(result)

	recieveChat := make(chan string, 4)

	User1 := User{ID: 101, Recieve: recieveChat}
	User2 := User{ID: 101, Recieve: recieveChat}

	go User1.Send("Hello, Good morning!")
	go User2.Send("Good morning!, How are you doing?")
	go User1.Send("Im cul, was there class today")
	go User2.Send("No not at all!")

	func() {
		for {
			time.Sleep(2 * time.Second)
			select {
			case msg := <-recieveChat:
				fmt.Println(msg)
			default:
				fmt.Println("Waiting for messages...")
			}
		}
	}()
}

type User struct {
	ID      int
	Recieve chan<- string
}

func (u User) Send(message string) {
	time.Sleep(time.Millisecond * 500)
	u.Recieve <- message
}
