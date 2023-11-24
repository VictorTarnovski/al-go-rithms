package main

import "fmt"

func ArraysAndSlices() {
	fmt.Println("Implementation of array")
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr)

	fmt.Println("Looping through the elements")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("Implementation of slices")

	// Go Slices
	// We can think of slices as simply slices of a cake
	// Slices does not copy the data into a new piece of memory, it only references the original array
	cake := [6]string{
		"dark-chocolate",
		"dark-chocolate",
		"dark-chocolate",
		"white-chocolate",
		"white-chocolate",
		"white-chocolate",
	}

	darkChocoalateSide := cake[0:3]
	whiteChocoalateSide := cake[3:6]

	fmt.Println("Entire cake before changes")
	fmt.Println(cake)

	// If we change the value of one element in the slice, it will reflect to the original array

	darkChocoalateSide[2] = "dark-chocolate-with-strawberry"
	whiteChocoalateSide[2] = "white-chocolate-with-coconut"
	fmt.Println("Entire cake after changes")
	fmt.Println(cake)

}
