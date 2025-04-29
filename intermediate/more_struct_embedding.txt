package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person // embed struct
	empId  string
	salary float64
}

// Method inheritance
func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

// Overriding method
func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f\n", e.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		person: person{
			name: "Juan",
			age:  30,
		},
		empId:  "E001",
		salary: 100000,
	}

	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	emp.introduce()
}
