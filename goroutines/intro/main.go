package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, I'm ", name)
}

func main() {
	greet("Divine")

	go greet("Osita")

	fmt.Println("main() routine continuing....")

	// the main goroutine has to wait sometime to allow 'Osita' go routine to execute Otherwise main() might exit before `Osita`

	time.Sleep(2 * time.Second)
	fmt.Println("Main go routine finished!")
}
