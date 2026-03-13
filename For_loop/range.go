package main

import "fmt"

func range3() {
	fmt.Println("Learning range")

	scores := []int{10, 20, 30, 40, 50}
	fmt.Println(scores)

	for index, value := range scores {
		fmt.Println(index, value)
	}
}
