package main

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"sort"
)

// --- slog ---

func demoSlog() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	logger.Info("slog demo", "feature", "structured logging", "version", "1.21")
}

// --- slices ---

func demoSlices() {
	nums := []int{3, 1, 4, 1, 5, 9}
	slices.Sort(nums)
	fmt.Printf("sorted: %v\n", nums)
	fmt.Printf("contains 4: %v\n", slices.Contains(nums, 4))
	fmt.Printf("index of 5: %d\n", slices.Index(nums, 5))
}

// --- maps ---

func demoMaps() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("keys: %v\n", keys)
	var vals []int
	for _, v := range m {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	fmt.Printf("values: %v\n", vals)
}

// --- builtins ---

func demoBuiltins() {
	fmt.Printf("min(3, 1, 4): %d\n", min(3, 1, 4))
	fmt.Printf("max(3, 1, 4): %d\n", max(3, 1, 4))
	m := map[string]int{"x": 1, "y": 2}
	clear(m)
	fmt.Printf("after clear: %v (len=%d)\n", m, len(m))
}

func main() {
	fmt.Println("=== Go 1.21 Feature Demos ===")
	fmt.Println()

	demoSlog()
	demoSlices()
	demoMaps()
	demoBuiltins()
}
