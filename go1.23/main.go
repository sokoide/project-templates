package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// --- Custom iterator using iter.Seq ---

func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

// --- Custom iterator using iter.Seq2 ---

func Enumerate[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

// --- Fibonacci iterator ---

func Fibonacci() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func demoCustomIterator() {
	fmt.Printf("countdown:")
	for v := range Countdown(5) {
		fmt.Printf(" %d", v)
	}
	fmt.Println()
}

func demoSeq2() {
	words := []string{"go", "is", "great"}
	fmt.Printf("enumerate:")
	for i, w := range Enumerate(words) {
		fmt.Printf(" [%d:%s]", i, w)
	}
	fmt.Println()
}

func demoFibonacci() {
	fmt.Printf("fibonacci (first 8):")
	count := 0
	for v := range Fibonacci() {
		fmt.Printf(" %d", v)
		count++
		if count >= 8 {
			break
		}
	}
	fmt.Println()
}

func demoStdlibIterators() {
	nums := []int{10, 20, 30}
	fmt.Printf("slices.All:")
	for i, v := range slices.All(nums) {
		fmt.Printf(" [%d:%d]", i, v)
	}
	fmt.Println()

	m := map[string]int{"x": 1, "y": 2}
	fmt.Printf("maps.All:")
	for k, v := range maps.All(m) {
		fmt.Printf(" %s=%d", k, v)
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Go 1.23 Feature Demos ===")
	fmt.Println()

	demoCustomIterator()
	demoSeq2()
	demoFibonacci()
	demoStdlibIterators()
}
