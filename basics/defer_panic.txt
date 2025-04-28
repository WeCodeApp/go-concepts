package main

import "fmt"

func main() {

	fmt.Println("Process 10")
	process(10)
	fmt.Println("Process -3")
	process(-3)
}

func process(input int) int {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		panic("input must be a non-negative integer")
		defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
	defer fmt.Println("Deferred 4")
	fmt.Println("Returning 8")
	return 8
}
