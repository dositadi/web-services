package main

import (
	"context"
	"fmt"
	"time"
)

func FetchDataFromDatabase(ctx context.Context, query string) {
	select {
	case <-time.After(4 * time.Second):
		fmt.Printf("%s: Successfully fetched data for query %s\n", time.Now().Format("15:04:05"), query)
	case <-ctx.Done():
		fmt.Printf("%s: Database query for %s timed out or was cancelled: %v\n", time.Now().Format("15:04:05"), query, ctx.Err())
		return
	}
}

func main() {
	// Create a parent context
	parentContext := context.Background()

	// Define a deadline
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(parentContext, timeout)
	defer cancel() // Always defer cancel() to free up querys later

	// Start a goroutine that performs tasks
	fmt.Printf("%s: Starting Database query with a timeout of %s.\n", time.Now().Format("15:04:05"), timeout)
	go FetchDataFromDatabase(ctx, "SELECT * FROM orders WHERE status = 'pending'")

	// Give goroutine time to run (and hit timeout)
	time.Sleep(3 * time.Second)

	// Example 2: Timeout longer than the task
	ctx2, cancel2 := context.WithTimeout(parentContext, 5*time.Second)
	defer cancel2()

	// start second goroutine
	fmt.Printf("%s: Starting another Database query with a longer timeout.\n", time.Now().Format("15:04:05"))
	go FetchDataFromDatabase(ctx2, "SELECT * FROM users")
	// Give goroutine time to run
	time.Sleep(6 * time.Second)
}
