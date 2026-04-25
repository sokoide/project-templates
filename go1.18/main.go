package main

import "fmt"

// --- Generic functions ---

func Map[T any, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](s []T, pred func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any, U any](s []T, init U, f func(U, T) U) U {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// --- Generic type ---

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

// --- Type constraints ---

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// --- Demo functions ---

func demoGenericFunctions() {
	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Printf("map: %v\n", doubled)

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Printf("filter: %v\n", evens)

	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("reduce: %d\n", sum)
}

func demoGenericType() {
	s := NewStack[string]()
	s.Push("hello")
	s.Push("world")
	v, _ := s.Pop()
	fmt.Printf("stack pop: %s (len=%d)\n", v, s.Len())
}

func demoTypeConstraint() {
	fmt.Printf("max(3, 7) = %d\n", Max(3, 7))
	fmt.Printf("max(\"abc\", \"xyz\") = %s\n", Max("abc", "xyz"))
}

func main() {
	fmt.Println("=== Go 1.18 Feature Demos ===")
	fmt.Println()

	demoGenericFunctions()
	demoGenericType()
	demoTypeConstraint()
}
