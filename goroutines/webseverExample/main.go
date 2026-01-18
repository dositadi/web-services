package main

import (
	"fmt"
	"time"
)

func simulateRequestProcessing(requestID int) {
	fmt.Println("Processing request, ", requestID, "...")
	time.Sleep(time.Duration(requestID) * time.Millisecond * 500)
	fmt.Println("Finished processing request, ", requestID)
}

func logRequest(requestID int) {
	fmt.Println("Logging request, ", requestID, "...")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Finished logging request, ", requestID)
}

func main() {
	fmt.Println("Sever Started...")

	for i := range 3 {
		requestID := i

		go simulateRequestProcessing(requestID)

		go logRequest(requestID)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("Sever ended.")
}
