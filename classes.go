package main

import "fmt"

func Classes() {
	type Person struct {
		name, surname string
	}

	// There are three ways to create a Instance of a Class in GO
	p1 := Person{name: "Victor", surname: "Tarnovski"}
	p2 := Person{"Victor", "Tarnovski"}
	p3 := new(Person)
	p3.name = "Victor"
	p3.surname = "Tarnovski"

	fmt.Println(p1, p2, p3)
}
