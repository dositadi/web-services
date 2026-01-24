package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var URLs []string = []string{
	"http://localhost:8081/word/1", "http://localhost:8081/roman/2", "http://localhost:8081/word/2", "http://localhost:8081/roman/1", "http://localhost:8081/foo/3", "http://localhost:8081/roman/4", "http://localhost:8081/word/-1", "http://localhost:8081/roman/3", "http://localhost:8081/word/5", "http://localhost:8081/game/6", "http://localhost:8081/word/6", "http://localhost:8081/roman/5", "http://localhost:8081/word/11", "http://localhost:8081/roman/8", "http://localhost:8081/word/8", "http://localhost:8081/roman/7", "http://localhost:8081/word/9", "http://localhost:8081/roman/10", "http://localhost:8081/word/10", "http://localhost:8081/roman/9",
}

func GetPath(urls []string) <-chan *http.Response {
	out := make(chan *http.Response)
	go func() {
		for _, url := range urls {
			response, err := http.Get(url)
			if err != nil {
				log.Fatal("Error making GET request: ", err)
			}
			out <- response
		}
		close(out)
	}()
	return out
}

func UnboxResponseBody(in <-chan *http.Response) <-chan string {
	out := make(chan string)

	go func() {
		for response := range in {
			defer response.Body.Close()
			responseBuffer, err := io.ReadAll(response.Body)
			if err != nil {
				out <- fmt.Errorf("Error reading response body: %w", err).Error()
			} else {
				out <- response.Status + "\n" + string(responseBuffer)
			}
		}
		close(out)
	}()
	return out
}

func main() {
	// Use pipeline to read the through the server
	response := GetPath(URLs)

	responseBody := UnboxResponseBody(response)

	for body := range responseBody {
		fmt.Println(body)
	}
}
