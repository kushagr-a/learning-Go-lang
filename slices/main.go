package main

import "fmt"

func main() {

	fmt.Println("Learning slices")
	// slices are dynamic arrays
	// common collection types
	// dynamic and can grow
	// []types{...}

	results := []string{"apple", "banana", "orange"}
	fmt.Println(results)

	// you access wise index

	// you can append empty 
	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20, 30)
	nums = append(nums, 30)
	fmt.Println(nums)
}