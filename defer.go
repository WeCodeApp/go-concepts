package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("1Defer in process", i)
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	i++
	fmt.Println("Normal execution")
	defer fmt.Println("2Defer in process", i)
}
