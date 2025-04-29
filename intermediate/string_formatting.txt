package main

import "fmt"

func main() {
	num := 424
	fmt.Printf("%05d\n", num)

	message := "Hello"
	fmt.Printf("|%10s|\n", message)  // right-aligned
	fmt.Printf("|%-10s|\n", message) // left-aligned

	message1 := "Hello \n World!"
	message2 := `Hello \nWorld!` // raw string literal
	// sqlQuery := `SELECT * FROM users WHERE name = 'O'Reilly'`

	fmt.Println(message1)
	fmt.Println(message2)

}
