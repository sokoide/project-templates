package main

import "testing"

func TestMap(t *testing.T) {
	result := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
	if len(result) != 3 || result[0] != 2 || result[1] != 4 || result[2] != 6 {
		t.Fatalf("unexpected map result: %v", result)
	}
}

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if len(result) != 2 || result[0] != 2 || result[1] != 4 {
		t.Fatalf("unexpected filter result: %v", result)
	}
}

func TestReduce(t *testing.T) {
	result := Reduce([]int{1, 2, 3}, 0, func(acc, n int) int { return acc + n })
	if result != 6 {
		t.Fatalf("expected 6, got %d", result)
	}
}

func TestStack(t *testing.T) {
	s := NewStack[int]()
	if s.Len() != 0 {
		t.Fatal("new stack should be empty")
	}
	s.Push(10)
	s.Push(20)
	if s.Len() != 2 {
		t.Fatalf("expected len 2, got %d", s.Len())
	}
	v, ok := s.Pop()
	if !ok || v != 20 {
		t.Fatalf("expected 20, got %d", v)
	}
	v, ok = s.Pop()
	if !ok || v != 10 {
		t.Fatalf("expected 10, got %d", v)
	}
	_, ok = s.Pop()
	if ok {
		t.Fatal("pop on empty stack should return false")
	}
}

func TestMax(t *testing.T) {
	if Max(3, 7) != 7 {
		t.Fatalf("expected 7")
	}
	if Max("abc", "xyz") != "xyz" {
		t.Fatalf("expected xyz")
	}
	if Max(3.14, 2.71) != 3.14 {
		t.Fatalf("expected 3.14")
	}
}
