package main

import "fmt"

var name2 = "Anna"

func main() {
	var age int
	var name string = "John"

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name2)

	count := 10

	printName()
	fmt.Println(count)
	// fmt.Println(lastName)

}

func printName() {
	lastName := "Smith"
	fmt.Println(lastName)
}
