package main

import "fmt"

/*
-> A variadic function in Go is a function that accepts a variable number
of arguments of the same type.
-> This is indicated by an ellipsis (...)
before the type of the last parameter in the function signature.
*/

func sumAll(nums ...int) int {
	total := 0

	for _, nums := range nums {
		total += nums
	}
	return total
}

func main() {
	fmt.Println("Varidatic function here we are larning")

	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumAll())

	// anomnous function with return value

	result := func(n int) int {
		return n * 2
	}
	fmt.Println("result of anomnous function", result(10))

	//IIFE (imidiatly invoked function expression)
	res1 := func(n int) int {
		return n * 4
	}(20)
	fmt.Println("res1 of IIFE", res1)

}
