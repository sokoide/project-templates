package main

import (
	"io"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestSetAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(1) // duplicate
	if len(s) != 2 {
		t.Fatalf("expected 2 members, got %d", len(s))
	}
}

func TestSetContains(t *testing.T) {
	s := NewSet[string]()
	s.Add("hello")
	if !s.Contains("hello") {
		t.Fatal("should contain hello")
	}
	if s.Contains("world") {
		t.Fatal("should not contain world")
	}
}

func TestSetMembers(t *testing.T) {
	s := NewSet[string]()
	s.Add("c")
	s.Add("a")
	s.Add("b")
	members := s.Members()
	slices.Sort(members)
	if !slices.Equal(members, []string{"a", "b", "c"}) {
		t.Fatalf("unexpected members: %v", members)
	}
}

func TestOsRootReadWrite(t *testing.T) {
	dir, err := os.MkdirTemp("", "go1.24-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	root, err := os.OpenRoot(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer root.Close()

	err = root.Mkdir("subdir", 0o755)
	if err != nil {
		t.Fatal(err)
	}

	f, err := root.Create("subdir/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString("test content"); err != nil {
		t.Fatal(err)
	}
	f.Close()

	f2, err := root.Open("subdir/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f2)
	f2.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "test content" {
		t.Fatalf("unexpected content: %s", data)
	}
}

func TestFilepathLocalize(t *testing.T) {
	result, err := filepath.Localize("safe/path.txt")
	if err != nil {
		t.Fatal(err)
	}
	if result != "safe/path.txt" {
		t.Fatalf("unexpected: %s", result)
	}
}
