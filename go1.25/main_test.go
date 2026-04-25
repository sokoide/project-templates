package main

import (
	"slices"
	"sync"
	"testing"
)

func TestConcurrentCollect(t *testing.T) {
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

	slices.Sort(collected)
	if !slices.Equal(collected, []int{0, 10, 20}) {
		t.Fatalf("unexpected: %v", collected)
	}
}

func TestPipeline(t *testing.T) {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}

	slices.Sort(result)
	if !slices.Equal(result, []int{2, 4, 6}) {
		t.Fatalf("unexpected: %v", result)
	}
}

func TestTimeout(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "response"

	var received string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		received = <-ch
	}()

	wg.Wait()

	if received != "response" {
		t.Fatalf("expected 'response', got '%s'", received)
	}
}
