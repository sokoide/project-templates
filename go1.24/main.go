package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

// --- Generic type alias ---

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Members() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}

// --- os.Root ---

func demoOsRoot() {
	dir, err := os.MkdirTemp("", "go1.24-demo")
	if err != nil {
		fmt.Printf("mkdir error: %v\n", err)
		return
	}
	defer os.RemoveAll(dir)

	root, err := os.OpenRoot(dir)
	if err != nil {
		fmt.Printf("openroot error: %v\n", err)
		return
	}
	defer root.Close()

	// Write via root
	f, err := root.Create("hello.txt")
	if err != nil {
		fmt.Printf("create error: %v\n", err)
		return
	}
	f.WriteString("hello from os.Root")
	f.Close()

	// Read via root
	f2, err := root.Open("hello.txt")
	if err != nil {
		fmt.Printf("open error: %v\n", err)
		return
	}
	data, _ := io.ReadAll(f2)
	f2.Close()
	fmt.Printf("os.Root read: %s\n", data)
}

func demoGenericTypeAlias() {
	s := NewSet[string]()
	s.Add("go")
	s.Add("rust")
	s.Add("go") // duplicate, ignored

	members := s.Members()
	sort.Strings(members)
	fmt.Printf("set members: %v\n", members)
	fmt.Printf("contains go: %v\n", s.Contains("go"))
	fmt.Printf("contains py: %v\n", s.Contains("py"))
}

// --- go.mod tool directive demo ---
// Go 1.24 allows tool dependencies in go.mod:
//   tool golang.org/x/tools/cmd/stringer
// This is a go.mod feature, shown here as documentation.

func demoToolDirective() {
	fmt.Println("go.mod tool directive: declare tool deps")
	fmt.Println("  e.g. tool golang.org/x/tools/cmd/stringer")
}

// --- filepath.Localize ---

func demoFilepathLocalize() {
	safe, err := filepath.Localize("subdir/file.txt")
	if err != nil {
		fmt.Printf("filepath.Localize error: %v\n", err)
		return
	}
	fmt.Printf("filepath.Localize: %s\n", safe)
}

func main() {
	fmt.Println("=== Go 1.24 Feature Demos ===")
	fmt.Println()

	demoGenericTypeAlias()
	demoOsRoot()
	demoToolDirective()
	demoFilepathLocalize()
}
