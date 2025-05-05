package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Command:", os.Args[0])
	//
	//for i, arg := range os.Args {
	//	fmt.Println("Argument", i, ":", arg)
	//}

	// define flags
	//var name string
	//var age int
	//var male bool
	//
	//flag.StringVar(&name, "name", "John", "Name of the user")
	//flag.IntVar(&age, "age", 18, "Age of the user")
	//flag.BoolVar(&male, "male", false, "Gender of the user")

	//flag.Parse()
	//fmt.Println("Name:", name)
	//fmt.Println("Age:", age)
	//fmt.Println("Male", male)

	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(stringFlag)

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Cmd processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of the request")
	flagsc2 := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires at least one argument.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCpmmand1")
		fmt.Println("Processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCpmmand2")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
