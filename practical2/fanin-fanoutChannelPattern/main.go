package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Image struct {
	ID       int
	Contents string
}

type Result struct {
	ImageID           int
	ProcessedContents string
	WorkerID          int
}

func ProcessImage(image Image, workerID int) Result {
	fmt.Printf("Worker %d: Processing Image %d...\n", workerID, image.ID)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	processedContent := fmt.Sprintf("Filtered(%s)_by_Worker%d\n", image.Contents, workerID)
	fmt.Printf("Worker %d: Finished Image %d...\n", workerID, image.ID)
	return Result{ImageID: image.ID, ProcessedContents: processedContent, WorkerID: workerID}
}

func main() {
	// Generate tasks (fan-out source)
	const numImages = 10
	images := make([]Image, numImages)
	for i := range numImages {
		images[i] = Image{ID: i + 1, Contents: fmt.Sprintf("Image%d.jpg", i+1)}
	}

	//create channels for input(fan-in) and output(fan-out)
	inputChan := make(chan Image, numImages)
	outputChan := make(chan Result, numImages)

	var wg sync.WaitGroup

	// Launch worker routines (fan-out)
	const numWorkers = 3
	fmt.Printf("Starting %d worker goroutines...\n", numWorkers)
	for i := range numWorkers {
		wg.Add(1)
		workerID := i + 1

		go func() {
			defer wg.Done()
			for img := range inputChan { // Worker reads from the shared input channel
				result := ProcessImage(img, workerID)
				outputChan <- result // Worker writes to shared output channel (Fan-out) process
			}
		}()
	}

	// send images to the input channel
	for _, img := range images {
		inputChan <- img
	}
	close(inputChan)

	// Wait for workers to finish and then close the output chan
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	//Collects results (fan-in)
	fmt.Println("\nCollecting results:")
	processedResults := make(map[int]Result)
	for result := range outputChan {
		fmt.Printf("Collected result for Image %d: %s (Processed by worker %d)\n", result.ImageID, result.ProcessedContents, result.WorkerID)
		processedResults[result.ImageID] = result
	}

	fmt.Printf("\nAll %d images processed. Total results collected: %d\n", numImages, len(processedResults))
}
