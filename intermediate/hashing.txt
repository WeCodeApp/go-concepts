package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	password := "password122"

	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	//fmt.Println("Password:", password)
	//fmt.Println("Hash256:", hash256)
	//fmt.Println("Hash512:", hash512)
	fmt.Printf("SHA-256 Hash hex val: %x:\n", hash256)
	fmt.Printf("SHA-512 Hash hex val: %x:\n", hash512)
}
