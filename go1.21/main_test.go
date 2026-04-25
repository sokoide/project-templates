package main

import (
	"slices"
	"testing"
)

func TestSlicesSort(t *testing.T) {
	nums := []int{3, 1, 2}
	slices.Sort(nums)
	if !slices.Equal(nums, []int{1, 2, 3}) {
		t.Fatalf("unexpected sorted: %v", nums)
	}
}

func TestSlicesContains(t *testing.T) {
	nums := []int{1, 2, 3}
	if !slices.Contains(nums, 2) {
		t.Fatal("should contain 2")
	}
	if slices.Contains(nums, 5) {
		t.Fatal("should not contain 5")
	}
}

func TestSlicesIndex(t *testing.T) {
	nums := []int{10, 20, 30}
	if slices.Index(nums, 20) != 1 {
		t.Fatalf("expected index 1, got %d", slices.Index(nums, 20))
	}
	if slices.Index(nums, 99) != -1 {
		t.Fatal("expected -1 for missing element")
	}
}

func TestMin(t *testing.T) {
	if min(3, 1, 4) != 1 {
		t.Fatalf("expected 1, got %d", min(3, 1, 4))
	}
}

func TestMax(t *testing.T) {
	if max(3, 1, 4) != 4 {
		t.Fatalf("expected 4, got %d", max(3, 1, 4))
	}
}

func TestClear(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	clear(m)
	if len(m) != 0 {
		t.Fatalf("expected empty map, got %d items", len(m))
	}
}
