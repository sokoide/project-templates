package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed static/hello.txt
var helloStr string

//go:embed static/hello.txt
var helloBytes []byte

//go:embed static
var staticFS embed.FS

func demoEmbedString() {
	fmt.Printf("embed string: %s", helloStr)
}

func demoEmbedBytes() {
	fmt.Printf("embed bytes: %s", helloBytes)
}

func demoEmbedFS() {
	data, err := fs.ReadFile(staticFS, "static/hello.txt")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("embed FS: %s", data)

	entries, _ := fs.ReadDir(staticFS, "static")
	fmt.Printf("embed FS entries:")
	for _, e := range entries {
		fmt.Printf(" %s", e.Name())
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Go 1.16 Feature Demos ===")
	fmt.Println()

	demoEmbedString()
	demoEmbedBytes()
	demoEmbedFS()
}
