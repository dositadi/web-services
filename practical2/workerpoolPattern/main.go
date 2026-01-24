package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// API rate limiting

type APIRequest struct {
	ID       int
	Endpoint string
}

type APIResponse struct {
	RequestID int
	Status    string
	Data      string
	Error     error
}

func APIWorker(id int, tasks <-chan APIRequest, results chan<- APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	for req := range tasks {
		fmt.Printf("Worker %d: Processing Request %d to %s\n", id, req.ID, req.Endpoint)

		time.Sleep(time.Duration(rand.Intn(500)+100) * time.Microsecond)

		// Simulate potential errors
		if rand.Intn(10) == 0 {
			results <- APIResponse{RequestID: req.ID, Status: "FAILED", Error: fmt.Errorf("Simulated API Error")}
			continue
		}

		results <- APIResponse{
			RequestID: req.ID,
			Status:    "SUCCESS",
			Data:      fmt.Sprintf("Data from %s for Request %d", req.Endpoint, req.ID),
		}
	}
}

func main() {
	const numRequests = 20
	const numWorkers = 5

	tasksChan := make(chan APIRequest, numRequests)
	responseChan := make(chan APIResponse, numRequests)

	var wg sync.WaitGroup

	// Set up worker functions
	for i := range numWorkers {
		wg.Add(1)
		go APIWorker(i+1, tasksChan, responseChan, &wg)
	}

	// Populate the task channel
	for i := range numRequests {
		tasksChan <- APIRequest{ID: i + 1, Endpoint: fmt.Sprintf("api/data/%d", i+1)}
	}
	close(tasksChan)

	go func() {
		wg.Wait()
		close(responseChan)
	}()

	// Collect and print response
	fmt.Println("\nCollecting API call results:")
	for res := range responseChan {
		if res.Error != nil {
			fmt.Printf("Request %d: %s (Error: %v)\n", res.RequestID, res.Status, res.Error)
		} else {
			fmt.Printf("Request %d: %s, %s\n", res.RequestID, res.Status, res.Data)
		}
	}

	fmt.Println("\nAll API requests processed.")
}
