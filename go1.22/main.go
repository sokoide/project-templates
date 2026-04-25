package main

import (
	"fmt"
	"net/http"
	"sync"
)

// --- Loop variable scope ---

func demoLoopVariable() {
	var wg sync.WaitGroup
	results := make([]int, 0, 3)
	var mu sync.Mutex

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			results = append(results, i) // Go 1.22: each gets own i
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("loop variable (goroutines): %v\n", results)
}

// --- HTTP routing ---

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Go 1.22: method + pattern matching
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter,
		r *http.Request) {
		fmt.Fprintf(w, "hello!")
	})

	// Go 1.22: path parameter {id}
	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter,
		r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "item: %s", id)
	})

	return mux
}

func demoHTTPRouting() {
	mux := setupMux()
	_ = mux // used in tests
	fmt.Println("HTTP routing: mux configured with " +
		"GET /hello and GET /items/{id}")
}

func main() {
	fmt.Println("=== Go 1.22 Feature Demos ===")
	fmt.Println()

	demoLoopVariable()
	demoHTTPRouting()
}
