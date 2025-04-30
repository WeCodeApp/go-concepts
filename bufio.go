package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(strings.NewReader("Hello World Golang! This is a senyence\nHow are you?"))
	//
	//data := make([]byte, 20)
	//n, err := reader.Read(data)
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	return
	//}
	//fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	//
	//line, err := reader.ReadString(',')
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	return
	//}
	//fmt.Println("Line read:", line)

	writer := bufio.NewWriter(os.Stdout)

	// Write byte s;ice
	data := []byte("Hello, bufio packagre!\n")

	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// writing string
	str := "This is a string. \n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing:", err)
		return
	}
}
