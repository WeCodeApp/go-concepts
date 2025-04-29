package main

import (
	"fmt"
)

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack elements: ")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	x1, y1 := 1, 2
	fmt.Println("Before swap:", x1, y1)
	x1, y1 = swap(x1, y1)
	fmt.Println("After swap:", x1, y1)

	x2, y2 := "Juan", "Pedro"
	fmt.Println("Before swap:", x2, y2)
	x2, y2 = swap(x2, y2)
	fmt.Println("After swap:", x2, y2)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	x, y := intStack.pop()
	fmt.Println("Popped:", x, y)
	fmt.Println("Is stack empty:", intStack.isEmpty())
	intStack.pop()
	fmt.Println("Popped:", x, y)
	intStack.pop()
	fmt.Println("Popped:", x, y)
	fmt.Println("Is stack empty:", intStack.isEmpty())
}
