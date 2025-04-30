package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	//password := "password123"

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
	}
	fmt.Println("Salt:", salt)
	fmt.Printf("Salt: %x\n", salt)

	// hash password with salt
	signUpHash := hashPassword("password123", salt)

	// Store the salt and pass in db
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("SaltStr:", saltStr)
	fmt.Println("signUpHash:", signUpHash)

	//verify password
	//retrieve the salt from db and decode
	decodeSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}
	loginHash := hashPassword("password124", decodeSalt)

	// compare the stored hash with the login hash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in")
	} else {
		fmt.Println("Login failed. Please check user credentials")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
