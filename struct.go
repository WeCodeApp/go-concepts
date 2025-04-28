package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func (p *Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		Address: Address{
			city:    "Quezon City",
			country: "PH",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "1231414",
			cell: "09123456789",
		},
	}
	fmt.Println("p1", p1)

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}
	fmt.Println("p2", p2)

	fmt.Println("p1", p1.firstName)
	fmt.Println("p2", p2.lastName)

	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Println("User:", user)

	fmt.Println("P1 full name:", p1.fullName())
	p1.incrementAgeByOne()
	fmt.Println("P1 age:", p1.age)

	fmt.Println("p1 address:", p1.Address)
	//p2.city = "Makati"
	//p2.country = "UK"

	fmt.Println("p2 address country:", p2.country)
	fmt.Println("p1 cell:", p1.cell)

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	fmt.Println("Are p1 and p2 equal", p1 == p2)
	fmt.Println("Are p3 and p2 equal", p3 == p2)
}
