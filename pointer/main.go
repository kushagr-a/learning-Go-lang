package main

import "fmt"

func main() {

	// store the memory address of a variable

	// &x -> address of x (make a pointer)
	// *p -> deref (go to that address and read / write)

	fmt.Println("learning pointer")
	score := 10
	fmt.Println("score berfore updating:", score)

	addScore(&score) // pass the address with the value
	// fmt.Println("score after updating:", score)

}

func addScore(score *int) {
	*score = *score + 10
	fmt.Println("score after updating:", *score)
}
