package main

import "fmt"

func main() {

	process(10)
	process(-3)
}

func process(input int) {
	if input < 0 {
		panic("input must be a non-negative integer")
	}
	fmt.Println("Processing input:", input)
}
