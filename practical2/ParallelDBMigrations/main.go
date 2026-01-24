package main

import (
	"fmt"
	"sync"
	"time"
)

func DataMigrations(id int, migrated chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Migrating data: %d....\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond)
	migrated <- fmt.Sprintf("%s: data%d migrated\n", time.Now().Format("15:39:01"), id)
	fmt.Printf("Migration of data %d done successfully.\n", id)
}

func main() {
	var wg sync.WaitGroup

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	migrated := make(chan string, len(data))

	for _, i := range data {
		wg.Add(1)
		go DataMigrations(i, migrated, &wg)
	}

	// Close channel once the data migration is over
	go func() {
		wg.Wait()
		close(migrated)
	}()

	for data := range migrated {
		fmt.Println(data)
	}
}
