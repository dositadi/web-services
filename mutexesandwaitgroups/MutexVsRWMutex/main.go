package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	fmt.Printf("Read: Key = %s, Value = %s\n", key, val)
	time.Sleep(5 * time.Millisecond)
	return val, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	fmt.Printf("Write: Key = %s, Value = %s\n", key, value)
	time.Sleep(10 * time.Millisecond)
}

func main() {
	cache := NewCache()

	var wg sync.WaitGroup
	wg.Add(7)

	// Writers
	go func() {
		defer wg.Done()
		cache.Set("name", "Alice")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		cache.Set("age", "30")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Millisecond)
		cache.Set("city", "New York")
	}()

	for i := 0; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			val, ok := cache.Get("name")
			if ok {
				fmt.Printf("Reader %d got name: %s\n", id, val)
			}
			val, ok = cache.Get("age")
			if ok {
				fmt.Printf("Reader %d got age: %s\n", id, val)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("\nFinal Cache state: %+v\n", cache.data)
}
