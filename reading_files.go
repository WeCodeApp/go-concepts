package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	//fmt.Println("File opened successfully")
	//
	//// read the contents of the opened file
	//data := make([]byte, 1024) // buffer to read data
	//_, err = file.Read(data)
	//if err != nil {
	//	fmt.Println("Error reading file", err)
	//	return
	//}
	//fmt.Println("File content:", string(data))

	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
}
