package main

import (
	"fmt"
)

func main() {
	data := []byte{'0'}
	fmt.Printf("data size: %v\n", len(data))

	content := []byte{'1', '1', '1', '1', '1'}
	fmt.Printf("content size: %v\n", len(content))

	data = append(data, content...)
	fmt.Printf("data size: %v\n", len(data))
	fmt.Printf("%v\n", data)
}
