package main

import "fmt"

func main() {
	// simple iteration
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over slice
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd numbers:", i)
		if i == 5 {
			break
		}
	}

	for i := range 10 {
		fmt.Println((10 - i))
	}
	fmt.Println("We have lift off!")

	rows := 5
	// Outer loop
	for i := 1; i <= rows; i++ {
		//inner loop
		for j := 1; j <= rows-1; j++ {
			fmt.Print(" ")
		}
		//inner loop
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
