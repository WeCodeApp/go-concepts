package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 40
	fmt.Println(numbers)

	numbers[0] = 90
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)
	fmt.Println("Third element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := &originalArray

	copiedArray[0] = 100
	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray) // dereferencing

	numbers1 := [5]int{25, 89, 43, 45, 100}

	for i := 0; i < len(numbers1); i++ {
		fmt.Println("Element at index", i, ":", numbers1[i])
	}

	for _, value := range numbers1 {
		fmt.Printf("Value: %d\n", value)
	}
}
