package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// write data to file
	data := []byte("Hello, world!\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file")

	file2, err := os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file2.Close()

	_, err = file2.WriteString("Hello Go!\n\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file")
}
