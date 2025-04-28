package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Array1 is equal to Array2", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

}
