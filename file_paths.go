package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	// Join paths using filepath join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println(joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println(normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Dir:", dir)
	fmt.Println("File:", file)
	fmt.Println("Base:", filepath.Base("/home/user/docs"))

	fmt.Println("Is relativePath variable absolute", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute", filepath.IsAbs(absolutePath))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("Rel:", rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("Rel:", rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Absolute path:", absPath)
	}
}
