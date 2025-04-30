package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println("Generates rand numbers 0-100:", rand.Intn(101))
	//fmt.Println("Generates rand numbers 5-10:", rand.Intn(6)+5)
	//
	//val := rand.New(rand.NewSource(time.Now().Unix()))
	//fmt.Println(val.Intn(10))

	for {
		// Show Menu
		fmt.Println("Welcome to Dice Game!")
		fmt.Println("1. Roll Dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter choice (1 or 2):")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice. Please enter 1 or 2.")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// show results
		fmt.Println("You rolled a", die1, "and a", die2)
		fmt.Println("Total:", die1+die2)

		// Ask of the user wants to roll again
		fmt.Println("Do you want to roll again? (y/n):")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}

		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}
