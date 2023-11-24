package main

import "fmt"

func Arrays() {
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
