package main

import (
	"fmt"
	"time"
)

func downloadFile(filename string, duration time.Duration) {
	fmt.Println("Starting download ", filename, "...")
	time.Sleep(duration) // Time to download
	fmt.Println("Finished download ", filename)
}

func main() {
	fmt.Println("Main routine started")

	go downloadFile("Gospel songs", 400*time.Millisecond)

	go downloadFile("Prayer chants", 200*time.Millisecond)

	go downloadFile("Gospel messages", 40*time.Millisecond)

	time.Sleep(3 * time.Second)

	fmt.Println("Main routine ended!")
}
