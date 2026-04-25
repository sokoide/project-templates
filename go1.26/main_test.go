package main

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestNewExpression(t *testing.T) {
	val := 42
	p := Person{Name: "Test", Age: new(val)}
	if p.Age == nil {
		t.Fatal("Age should not be nil")
	}
	if *p.Age != 42 {
		t.Fatalf("expected 42, got %d", *p.Age)
	}
}

func TestSelfReferential(t *testing.T) {
	result := algo(Int(10), Int(20))
	if result != 30 {
		t.Fatalf("expected 30, got %d", result)
	}
}

func TestAsTypeMatch(t *testing.T) {
	err := findItem("test")
	ne, ok := errors.AsType[*NotFoundError](err)
	if !ok {
		t.Fatal("expected NotFoundError")
	}
	if ne.Name != "test" {
		t.Fatalf("expected 'test', got '%s'", ne.Name)
	}
}

func TestAsTypeNoMatch(t *testing.T) {
	err := errors.New("some error")
	_, ok := errors.AsType[*NotFoundError](err)
	if ok {
		t.Fatal("expected no match for non-matching error")
	}
}

func TestReflectFieldsIterator(t *testing.T) {
	p := Point{X: 1, Y: 2}
	v := reflect.ValueOf(p)
	count := 0
	for field, val := range v.Fields() {
		count++
		if field.Name == "X" && val.Int() != 1 {
			t.Fatalf("X should be 1, got %d", val.Int())
		}
		if field.Name == "Y" && val.Int() != 2 {
			t.Fatalf("Y should be 2, got %d", val.Int())
		}
	}
	if count != 2 {
		t.Fatalf("expected 2 fields, got %d", count)
	}
}

func TestReflectMethodsIterator(t *testing.T) {
	v := reflect.ValueOf(Int(0))
	count := 0
	for range v.Methods() {
		count++
	}
	if count == 0 {
		t.Fatal("Int should have at least the Add method")
	}
}

func TestBufferPeek(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("hello")
	peeked, _ := buf.Peek(3)
	if string(peeked) != "hel" {
		t.Fatalf("expected 'hel', got '%s'", peeked)
	}
	// Peek should not consume bytes
	if buf.Len() != 5 {
		t.Fatalf("buffer should still have 5 bytes, got %d", buf.Len())
	}
}
