package main

import "fmt"

func main() {
	fmt.Println("defer in go")

	// defer -> delay execution until the surrounding function returns

	// LIFO (Last In First Out)

	// The defer statement in Go is used to postpone the execution of a function 
	// until the surrounding function returns

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	fmt.Println("hello")

	// output: hello three two one

	// defer is often used for cleanup
	// file.Close()
	// db.Close()
	// network.Close()

	// example

}
