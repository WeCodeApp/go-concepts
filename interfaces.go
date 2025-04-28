package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Properties:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perim:", g.perim())
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 7}

	fmt.Println("Measuring Rectangle")
	measure(r)

	fmt.Println("Measuring Circle")
	measure(c)

	myPrinter(1, "Juan", 3.14, true)
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println("My printer", v)
	}
}
