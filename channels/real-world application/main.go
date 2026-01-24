package main

import (
	"fmt"
	"net/http"
	"time"
)

// Request-Response Handling in a Web Server

type Request struct {
	ID       string
	URL      string
	Response chan<- string // channel to recieve response!
}

func PocessRequest(id, url string) string {
	fmt.Println("Processing request: ", id, " with URL: ", url)
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("Response for %v:%v processed", id, url)
}

func Response(requests <-chan Request) {
	for req := range requests {
		responseMessage := PocessRequest(req.ID, req.URL)
		req.Response <- responseMessage
	}
}

func main() {
	requests := make(chan Request)
	requestsRange := 5

	for i := 0; i < requestsRange; i++ {
		go Response(requests)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responseChannel := make(chan string)
		reqID := fmt.Sprintf("%d", time.Now().UnixNano())

		requests <- Request{ID: reqID, URL: r.URL.Path, Response: responseChannel}

		response := <-responseChannel
		fmt.Fprintf(w, response)

		fmt.Println(r.URL.Path)
		close(responseChannel)
	})

	fmt.Println("Sever starting on :8080")
	http.ListenAndServe(":8080", nil)
}
