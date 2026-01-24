package main

import (
	"fmt"
	"sync"
	"time"
)

type URLFetcher struct {
	URL map[string]int
	mu  sync.Mutex
}

func Fetch(url string, wg *sync.WaitGroup, fetcher *URLFetcher) int {
	defer wg.Done()
	fetcher.mu.Lock()
	defer fetcher.mu.Unlock()

	time.Sleep(500 * time.Millisecond)

	return fetcher.URL[url]
}

func main() {
	URLs := map[string]int{
		"go.dev":     100,
		"python.dev": 20,
		"java.dev":   50,
		"unity.dev":  605,
		"engage.com": 122,
	}

	urlFectcher := URLFetcher{
		URL: URLs,
	}

	var wg sync.WaitGroup

	wg.Add(5)

	i := 0

	for key := range URLs {
		go func(id int) {
			fmt.Printf("Goroutine %d running\n", id)
			contentSize := Fetch(key, &wg, &urlFectcher)
			fmt.Printf("The size of '%s' content is %v\n", key, contentSize)
		}(i)
		i++
	}

	fmt.Println("Main goroutine: Awaiting URL's content fetch!...")
	wg.Wait()
	fmt.Println("Main goroutine: Fetched URL's contents successfully!")
}
