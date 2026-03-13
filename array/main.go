package main

import "fmt"

func main() {

	// we are going to work on slices cause slices are dynamic arrays
	// array is fixed size

	// fixed and can not grow or shrink
	var marks [3]int = [3]int{90, 80, 70}
	fmt.Println(marks)

	// marks := [3]int{90,80,70}

	// array literal
	res := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(res))
}
									