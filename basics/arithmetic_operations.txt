package main

import (
	"fmt"
	"math"
)

func main() {
	var maxInt int64 = 9223372036854775807

	fmt.Println("Maximum signed integer:", maxInt)
	maxInt = maxInt + 1
	fmt.Println("Maximum signed integer overflow:", maxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println("Small float:", smallFloat)
	smallFloat = smallFloat / math.MaxFloat32
	fmt.Println("Small float:", smallFloat)

	const P float64 = 22 / 7.0
	fmt.Println("Result:", P)

}
