package main

import (
	"fmt"
	"sync"
	"time"
)

// --- Go 1.25: testing/synctest ---
// synctest provides deterministic testing of concurrent code.
// The demo functions below show patterns that synctest enables.

func demoConcurrentCollect() {
	var wg sync.WaitGroup
	results := make(chan int, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * 10
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var collected []int
	for v := range results {
		collected = append(collected, v)
	}
	fmt.Printf("concurrent collect: %v\n", collected)
}

func demoPipeline() {
	in := make(chan int)
	out := make(chan int)

	// Stage 1: generate numbers
	go func() {
		for i := 1; i <= 3; i++ {
			in <- i
		}
		close(in)
	}()

	// Stage 2: double
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()

	// Collect
	var result []int
	for v := range out {
		result = append(result, v)
	}
	fmt.Printf("pipeline result: %v\n", result)
}

func demoTimeout() {
	ch := make(chan string, 1)
	ch <- "quick response"

	select {
	case msg := <-ch:
		fmt.Printf("timeout demo: %s\n", msg)
	case <-time.After(time.Second):
		fmt.Println("timeout demo: timed out")
	}
}

func main() {
	fmt.Println("=== Go 1.25 Feature Demos ===")
	fmt.Println()

	demoConcurrentCollect()
	demoPipeline()
	demoTimeout()
}
