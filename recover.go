package main

import "fmt"

func main() {
	process()
	fmt.Println("Returned from process")
}

func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start process")
	panic("Something went wrong")
	fmt.Println("End process")
}
