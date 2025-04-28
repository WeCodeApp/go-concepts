package main

import "fmt"

func main() {
	statement, total := sum("The total is", 1, 2, 3, 4, 5)
	fmt.Println(statement, total)
}

func sum(returnString string, nums ...int) (string, int) {
	total := 0

	for _, v := range nums {
		total += v
	}

	return returnString, total
}
