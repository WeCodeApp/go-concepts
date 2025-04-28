package main

import "fmt"

func main() {
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's an banana")
	default:
		fmt.Println("Unknown fruit")
	}

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Satuday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Iinvalid day")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	num := 0
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not 2")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

func checkType(x any) {
	switch x.(type) {
	case int:
		fmt.Println("Its an integer")
	case float64:
		fmt.Println("Its a float64")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Unknown type")
	}
}
