package main

import (
	"fmt"
	"strings"
)

func main() {
	// String builder
	var builder strings.Builder

	// Write strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// Convert builder to a string
	result := builder.String()
	fmt.Println("1:", result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you?")

	result = builder.String()
	fmt.Println("2:", result)

	// Reset builder
	builder.Reset()
	builder.WriteString("Start fresh")
	result = builder.String()
	fmt.Println("3:", result)
}
