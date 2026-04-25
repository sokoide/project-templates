package main

import (
	"iter"
	"slices"
	"testing"
)

func TestCountdown(t *testing.T) {
	var result []int
	for v := range Countdown(3) {
		result = append(result, v)
	}
	if !slices.Equal(result, []int{3, 2, 1}) {
		t.Fatalf("unexpected countdown: %v", result)
	}
}

func TestCountdownBreak(t *testing.T) {
	var result []int
	for v := range Countdown(10) {
		result = append(result, v)
		if v == 7 {
			break
		}
	}
	if !slices.Equal(result, []int{10, 9, 8, 7}) {
		t.Fatalf("unexpected break result: %v", result)
	}
}

func TestEnumerate(t *testing.T) {
	words := []string{"a", "b", "c"}
	var indices []int
	var values []string
	for i, v := range Enumerate(words) {
		indices = append(indices, i)
		values = append(values, v)
	}
	if !slices.Equal(indices, []int{0, 1, 2}) {
		t.Fatalf("unexpected indices: %v", indices)
	}
	if !slices.Equal(values, []string{"a", "b", "c"}) {
		t.Fatalf("unexpected values: %v", values)
	}
}

func TestFibonacci(t *testing.T) {
	var result []int
	count := 0
	for v := range Fibonacci() {
		result = append(result, v)
		count++
		if count >= 6 {
			break
		}
	}
	if !slices.Equal(result, []int{0, 1, 1, 2, 3, 5}) {
		t.Fatalf("unexpected fibonacci: %v", result)
	}
}

func TestIterSeqType(t *testing.T) {
	var _ iter.Seq[int] = Countdown(1)
	var _ iter.Seq2[int, string] = Enumerate([]string{"x"})
}
