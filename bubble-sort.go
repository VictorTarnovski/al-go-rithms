package main

import "fmt"

func BubbleSort() {
	list := [5]int{9, 5, 7, 1, 3}

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}

	fmt.Println(list)
}
