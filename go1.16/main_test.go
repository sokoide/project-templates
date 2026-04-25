package main

import (
	"io/fs"
	"testing"
)

func TestEmbedString(t *testing.T) {
	if helloStr != "Hello from embedded file!\n" {
		t.Fatalf("unexpected embed string: %q", helloStr)
	}
}

func TestEmbedBytes(t *testing.T) {
	expected := "Hello from embedded file!\n"
	if string(helloBytes) != expected {
		t.Fatalf("unexpected embed bytes: %q", string(helloBytes))
	}
}

func TestEmbedFS(t *testing.T) {
	data, err := fs.ReadFile(staticFS, "static/hello.txt")
	if err != nil {
		t.Fatalf("read error: %v", err)
	}
	if string(data) != "Hello from embedded file!\n" {
		t.Fatalf("unexpected FS content: %q", string(data))
	}
}

func TestEmbedFSEntries(t *testing.T) {
	entries, err := fs.ReadDir(staticFS, "static")
	if err != nil {
		t.Fatalf("readdir error: %v", err)
	}
	names := make(map[string]bool)
	for _, e := range entries {
		names[e.Name()] = true
	}
	if !names["hello.txt"] {
		t.Fatal("expected hello.txt in embedded FS")
	}
	if !names["data.json"] {
		t.Fatal("expected data.json in embedded FS")
	}
}
