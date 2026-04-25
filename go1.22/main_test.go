package main

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"testing"
)

func TestLoopVariableScope(t *testing.T) {
	var wg sync.WaitGroup
	results := make([]int, 0, 3)
	var mu sync.Mutex

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			results = append(results, i)
			mu.Unlock()
		}()
	}
	wg.Wait()

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	slices.Sort(results)
	for i, v := range results {
		if v != i {
			t.Fatalf("expected %d at index %d, got %d", i, i, v)
		}
	}
}

func TestHTTPHello(t *testing.T) {
	mux := setupMux()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "hello!" {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestHTTPItemWithPathParam(t *testing.T) {
	mux := setupMux()
	req := httptest.NewRequest(http.MethodGet, "/items/42", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "item: 42" {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestHTTPMethodNotAllowed(t *testing.T) {
	mux := setupMux()
	req := httptest.NewRequest(http.MethodPost, "/hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Fatal("POST should not match GET pattern")
	}
}
