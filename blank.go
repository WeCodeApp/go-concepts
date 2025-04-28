package main

import "fmt"

func main() {
	a, _ := someFunction()
	fmt.Println(a)
}

func someFunction() (int, int) {
	return 1, 2
}
