package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"time"
)

// --- Go 1.26: new with expression ---

type Person struct {
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

func yearsSince(t time.Time) int {
	return int(time.Since(t).Hours() / (365.25 * 24))
}

func demoNewExpression() {
	born := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	p := Person{
		Name: "Alice",
		Age:  new(yearsSince(born)), // Go 1.26: new with expression
	}
	data, _ := json.Marshal(p)
	fmt.Printf("new expression: %s\n", data)
}

// --- Go 1.26: self-referential type constraints ---

type Adder[A Adder[A]] interface {
	Add(A) A
}

type Int int

func (a Int) Add(b Int) Int { return a + b }

func algo[A Adder[A]](x, y A) A {
	return x.Add(y)
}

func demoSelfReferential() {
	result := algo(Int(3), Int(4))
	fmt.Printf("self-referential: algo(3, 4) = %d\n", result)
}

// --- Go 1.26: errors.AsType ---

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return "not found: " + e.Name }

func findItem(name string) error {
	return &NotFoundError{Name: name}
}

func demoAsType() {
	err := findItem("widget")
	if ne, ok := errors.AsType[*NotFoundError](err); ok {
		fmt.Printf("AsType: %s\n", ne.Error())
	}
}

// --- Go 1.26: reflect iterators ---

type Point struct {
	X int
	Y int
}

func demoReflectIterators() {
	p := Point{X: 10, Y: 20}
	t := reflect.TypeOf(p)
	fmt.Printf("reflect fields:")
	for field := range t.Fields() {
		fmt.Printf(" %s", field.Name)
	}
	fmt.Println()

	v := reflect.ValueOf(p)
	fmt.Printf("reflect values:")
	for field, val := range v.Fields() {
		fmt.Printf(" %s=%v", field.Name, val.Int())
	}
	fmt.Println()
}

// --- Go 1.26: slog.NewMultiHandler ---

func demoMultiHandler() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	handler := slog.NewMultiHandler(textHandler)
	logger := slog.New(handler)
	logger.Info("multi-handler demo", "key", "value")
}

// --- Go 1.26: bytes.Buffer.Peek ---

func demoBufferPeek() {
	var buf bytes.Buffer
	buf.WriteString("hello world")
	peeked, _ := buf.Peek(5)
	fmt.Printf("peek: %s (buf still has %d bytes)\n", peeked, buf.Len())
}

func main() {
	fmt.Println("=== Go 1.26 Feature Demos ===")
	fmt.Println()

	demoNewExpression()
	demoSelfReferential()
	demoAsType()
	demoReflectIterators()
	demoMultiHandler()
	demoBufferPeek()
}
