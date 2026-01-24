package main

import (
	"fmt"
	"sync"
	"time"
)

// Demonstrating example without mutex
func exampleWithoutMutex() {
	var counter int
	numGoRoutines := 1000

	var wg sync.WaitGroup

	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			defer wg.Done()
			// Each Go routine increments the counter a 100 times
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Counter without mutex: %d (Expected: %d)\n", counter, numGoRoutines*100)
}

// Demonstrating example with mutex
func ExampleWithMutex() {
	numOfGoRoutines := 1000
	var counter int
	var mu sync.Mutex

	var wg sync.WaitGroup

	wg.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock() // Lock the shared resource once a goroutine is making use of it, so that other go routines will have to wait
				counter++
				mu.Unlock() // Unlock the shared resource once a goroutine is done making use of it, so that the next go routines can make use of it
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Counter with mutex: %d (Expected: %d)\n", counter, numOfGoRoutines*100)
}

// A more complex example: Managing a shared inventory with a mutex
type Inventory struct {
	items map[string]int
	mu    sync.Mutex
}

func (inv *Inventory) AddItem(item string, qty int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	inv.items[item] += qty
	fmt.Printf("Added %v of %s. Current stock: %v\n", qty, item, inv.items[item])
	time.Sleep(10 * time.Millisecond)
}

func (inv *Inventory) RemoveItem(item string, qty int) bool {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if inv.items[item] >= qty {
		inv.items[item] -= qty
		fmt.Printf("Removed %s of %v. Current stock: %v\n", item, qty, inv.items[item])
		time.Sleep(10 * time.Millisecond)
		return true
	}
	fmt.Printf("Not enough item: %s, to remove: %v. Current stock: %v\n", item, qty, inv.items[item])
	return false
}

func (inv *Inventory) GetStock(item string) int {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	return inv.items[item]
}

func main() {
	fmt.Println("---Demonstrating race condition without mutex---")
	exampleWithoutMutex()
	fmt.Println("---Demonstrating race condition with mutex---")
	ExampleWithMutex()

	// Initialize inventory
	storeInventory := &Inventory{
		items: make(map[string]int),
	}

	storeInventory.AddItem("Laptop", 10)

	var wg sync.WaitGroup

	wg.Add(5)

	go func() {
		defer wg.Done()
		storeInventory.AddItem("Mouse", 5)
		storeInventory.RemoveItem("Laptop", 2)
	}()

	go func() {
		defer wg.Done()
		storeInventory.RemoveItem("Laptop", 8)
		storeInventory.AddItem("Keyboard", 3)
	}()

	go func() {
		defer wg.Done()
		storeInventory.AddItem("Laptop", 5)
	}()

	go func() {
		defer wg.Done()
		storeInventory.RemoveItem("Mouse", 3)
	}()

	go func() {
		defer wg.Done()
		stock := storeInventory.GetStock("Laptop")
		fmt.Printf("Initial Laptop stock check: %v\n", stock)
		time.Sleep(10 * time.Millisecond)
		stock = storeInventory.GetStock("Laptop")
		fmt.Printf("Final Laptop stock check: %v\n", stock)
	}()
	wg.Wait()
	fmt.Printf("\nFinal Inventory Status: %+v\n", storeInventory.items)
}
