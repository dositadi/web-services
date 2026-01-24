package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type contextKey string

const userKey contextKey = "user"
const traceIDKey contextKey = "traceID"

// Mock auth middleware
func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentiaction
		userID := "user-123"
		newCtx := context.WithValue(r.Context(), userKey, userID)
		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentiaction
		traceID := fmt.Sprintf("trace-%d", time.Now().UnixNano())
		newCtx := context.WithValue(r.Context(), traceIDKey, traceID)
		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the requests context
	userID, ok := r.Context().Value(userKey).(string)
	if !ok {
		userID = "anonymous"
	}

	traceID, ok := r.Context().Value(traceIDKey).(string)
	if !ok {
		traceID = "N/A"
	}

	fmt.Printf("Trace ID: %s, Handling request for user: %s\n", traceID, userID)
	fmt.Fprint(w, "Hello, "+userID+", Trace ID: "+traceID+"\n")
}

func main() {
	//Set up a simple HTTP server with middleware
	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleWare(authMiddleWare(http.HandlerFunc(handleRequest))))

	fmt.Println("Server listening on :8080")
	//To test open http://localhost:8080/hello
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
