package main

import "fmt"

// In GO conventional classes doesn't exist
// GO "classes" are simply struct's that therefore we implement functions for those structs

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func Classes() {
	// We can create an instance of a struct in many ways
	// By declaring a variable and then values to its properties
	var p1 Person
	p1.name = "John Doe"
	p1.age = 30

	// By declaring the values right on declaration
	p2 := Person{
		name: "Sue Widget",
		age:  35,
	}

	// Same as before but with only the values
	p3 := Person{"Dominic Down", 27}

	fmt.Println(p1, p2, p3)

	p1.SetAge(1)

	fmt.Println(p1, p2, p3)

}
