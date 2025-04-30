package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempFileName := "temporaryFile"

	tempfile, err := os.CreateTemp("", tempFileName)
	checkError(err)

	fmt.Println("Temporary file created:", tempfile.Name())

	defer os.Remove(tempfile.Name())
	defer tempfile.Close()

	tempDir, err := os.MkdirTemp("", "tempDir")
	checkError(err)

	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary directory created:", tempDir)
}
