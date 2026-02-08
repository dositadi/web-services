package main

import (
	"fmt"
	"sync"
	"time"
)

func DoingWork(id int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	fmt.Printf("Registering user %v...\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Finished Registering user %v.\n", id)
}

func Download(id int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	fmt.Printf("Downloading img%v.jpeg...\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("img%v.jpeg Downloaded completely.\n", id)
}

func ReadSlice(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range len(numbers) {
			out <- numbers[i]
		}
		close(out)
	}()
	return out
}

func SquareInt(in <-chan int) <-chan int {
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
	fmt.Println("Main Routine: starting....")
	var mu sync.Mutex
	var wg sync.WaitGroup

	func() {
		for i := range 5 {
			wg.Add(2)
			go DoingWork(i, &mu, &wg)
			go Download(i, &mu, &wg)
		}
	}()

	wg.Wait()
	fmt.Println("")

	slice := []int{1, 2, 3, 4, 5, 8}

	values := ReadSlice(slice)

	squares := SquareInt(values)

	for s := range squares {
		fmt.Println(s)
	}

	out := make(chan int, 5)
	for i := range 4 {
		out <- i
	}
	out <- 5
	i := <-out
	fmt.Println(cap(out), ", ", len(out))
	fmt.Println(i)
	fmt.Println()
	out <- 6

	for i := range 5 {
		fmt.Println(<-out, i)
	}
}
