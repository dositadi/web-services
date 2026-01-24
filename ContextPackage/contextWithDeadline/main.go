package main

import (
	"context"
	"fmt"
	"time"
)

func NetworkRequest(ctx context.Context, resource string) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("%s: Successfully fetched %s\n", time.Now().Format("15:04:05"), resource)
	case <-ctx.Done():
		fmt.Printf("%s: Network request for %s failed due to: %v\n", time.Now().Format("15:04:05"), resource, ctx.Err())
		return
	}
}

func main() {
	// Create a parent context
	parentContext := context.Background()

	// Define a deadline
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(parentContext, deadline)
	defer cancel() // Always defer cancel() to free up resources later

	// Start a goroutine that performs tasks
	fmt.Printf("%s: Starting Network request with a deadline of %s.\n", time.Now().Format("15:04:05"), deadline.Format("15:04:05"))
	go NetworkRequest(ctx, "user_profile")

	// Give goroutine time to run and hit deadline
	time.Sleep(2 * time.Second)

	// Example 2
	ctx2, cancel2 := context.WithDeadline(parentContext, time.Now().Add(5*time.Second))
	defer cancel2()

	// start second goroutine
	fmt.Printf("%s: Starting another Network request with a longer deadline.\n", time.Now().Format("15:04:05"))
	go NetworkRequest(ctx2, "product_list")
	// Give goroutine time to run
	time.Sleep(4 * time.Second)
}
