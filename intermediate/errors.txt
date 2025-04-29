package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("math: square root of negative number: %v", x)
	}
	// compute square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	// process data
	return nil
}

func main() {
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Square root:", result)

	//result1, err1 := sqrt(-16)
	//if err1 != nil {
	//	fmt.Println(err1)
	//	return
	//}
	//fmt.Println("Square root:", result1)

	data := []byte{}
	err2 := process(data)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println("Data processed successfully")
}
