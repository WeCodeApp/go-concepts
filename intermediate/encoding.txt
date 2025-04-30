package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 Encoding")

	// encode to base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Base64 encoded string", encoded)

	// decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded string", string(decoded))

	// URL safe, avoid + and /
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL safe encoded string", urlSafeEncoded)

}
