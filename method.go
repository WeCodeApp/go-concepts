package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Printf("Area of a rectangle with width %v and length %v is %v\n", rect.width, rect.length, area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Printf("Area of a rectangle with width %v and length %v is %v\n", rect.width, rect.length, area)

}
