package main

import (
	"fmt"
)

/*
Generics, introduced in Go 1.18, enable developers to write reusable,
type-safe code that works with functions and data structures
across a variety of types without sacrificing compile-time safety
*/

// func printSlice [T comparable] (items []T) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }

// here only allow int and string type
func printSlice[T int | string](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

// func printStringSlice(items []string) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }

// using struct

// lifo
type Stack[T int | string] struct {
	element []T
}

func main() {
	fmt.Println("learning generics")

	// names := []string{"a", "b", "c", "d", "e"}
	// printSlice(names)

	// fmt.Println("\n")

	// nums := []int{1, 2, 3, 4, 5}
	// printSlice(nums)
	// printStringSlice(names)

	mystack := Stack[int]{
		element: []int{1, 2, 3, 4, 5},
		// element: []string{"a", "b", "c", "d", "e"},
	}

	fmt.Println(mystack)
}
