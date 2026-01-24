package main

import (
	"fmt"
	"sync"
	"time"
)

func ProcessTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("GoRoutine %d: Starting task...\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("GoRoutine %d: Task completed.\n", id)
}

func main() {
	var wg sync.WaitGroup

	numTasks := 5
	wg.Add(numTasks)

	fmt.Println("Launching goroutines...")
	for i := 1; i <= numTasks; i++ {
		go ProcessTask(i, &wg)
	}

	fmt.Println("main GoRoutine waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("Main GoRoutine: All tasks completed. Exiting.")

	// Example scenerio downloading multiple files concurrently
	fmt.Println("\n--- Concurrent File Download Simulation ---")
	filesToDownload := []string{"image.jpg", "document.pdf", "archive.zip", "video.mp4"}
	var downloadingWG sync.WaitGroup
	downloadingWG.Add(len(filesToDownload))

	for i, file := range filesToDownload {
		go func(taskID int, filename string) {
			defer downloadingWG.Done()
			fmt.Printf("Downloading '%s' (Task %d)...\n", filename, taskID)
			time.Sleep(time.Duration(taskID+1) * 200 * time.Millisecond)
			fmt.Printf("Finished downloading '%s' (Task %d).\n", filename, taskID)
		}(i+1, file)
	}
	fmt.Println("Main goroutine: Waiting for all downloads to complete...")
	downloadingWG.Wait()
	fmt.Println("Main goroutine: All downloads completed. Proceeding to process files.")

	// Example scenerio: Processing a batch of items
	fmt.Println("\n--- Batch Processing Simulation ---")
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}

	var batchWG sync.WaitGroup
	batchSize := 2
	processedItems := make(chan int, len(items))

	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}
		batch := items[i:end]

		batchWG.Add(1)
		go func(batchSubset []int) {
			defer batchWG.Done()
			fmt.Printf("Processing batch: %v\n", batchSubset)
			for _, item := range batchSubset {
				time.Sleep(50 * time.Millisecond)
				processedItems <- item * 2
			}
		}(batch)
	}
	// Start goroutine to close the channel after all the batches are done
	go func() {
		batchWG.Wait()
		close(processedItems)
	}()

	fmt.Println("Main goroutine: Collecting processed results...")
	for result := range processedItems {
		fmt.Printf("Collected result: %d\n", result)
	}
	fmt.Println("All batch processing results collected.")
}
