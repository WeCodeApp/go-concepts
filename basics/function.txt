package main

import "fmt"

func main() {
	//fmt.Println("Sum", add(1, 2))
	//
	//greet := func() {
	//	fmt.Println("Hello anonymous function")
	//}
	//
	//greet()
	//
	//operation := add
	//result := operation(6, 4)
	//fmt.Println("Result", result)

	// pass a function as an argument
	result := applyOperation(2, 3, add)
	fmt.Println("2+3", result)

	// return and using a funtion
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6*2", multiplyBy2(6))
	fmt.Println("5*2", multiplyBy2(5))
}

func add(a, b int) int {
	return a + b
}

// function that takes a function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
