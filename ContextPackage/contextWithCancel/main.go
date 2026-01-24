package main

import (
	"context"
	"fmt"
	"time"
)

func PerformTask(ctx context.Context, taskname string) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("%s: Task %s completed successfully.\n", time.Now().Format("15:04:05"), taskname)
	case <-ctx.Done():
		fmt.Printf("%s: Task %s cancelled: %v\n", time.Now().Format("15:04:05"), taskname, ctx.Err())
	}
}

func main() {
	// Create a parent context
	parentContext := context.Background()

	// Derive a cancellable context from the parent
	ctx, cancel := context.WithCancel(parentContext)

	// Start a goroutine that performs tasks
	go PerformTask(ctx, "Long running task 1")

	// Wait for a short duration and then cancel the context
	time.Sleep(1 * time.Second)
	fmt.Printf("%s: Main routine cancelling context.\n", time.Now().Format("15:04:05"))
	cancel()

	// Give some time for main to react to cancellation
	time.Sleep(500 * time.Millisecond)

	ctx2, cancel2 := context.WithCancel(parentContext)

	// start second goroutine
	go PerformTask(ctx2, "Another task running")

	// Wait for a short duration and then cancel the context
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("%s: Main routine cancelling another context.\n", time.Now().Format("15:04:05"))
	cancel2()

	time.Sleep(500 * time.Millisecond)
}
