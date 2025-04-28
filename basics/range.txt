package main

import "fmt"

func main() {
	message := "Hello, World!"

	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Value: %c\n", i, v)
	}
}
