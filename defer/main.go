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

	fmt.Println("case 1: success")
	if err := dowork(true); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("case 2: failure")
	if err := dowork(false); err != nil {
		fmt.Println("error:", err)
	}

	// output: hello three two one

	// defer is often used for cleanup
	// file.Close()
	// db.Close()
	// network.Close()

	// example

}

func dowork(success bool) error {
	// resorce related
	// start message -> resource acquired
	// clenup

	fmt.Println("start: resource aquire")

	// defer will guarenteed to run even if function returns early
	// success or failure
	defer fmt.Println("cleanup: resource released")

	if !success {
		return fmt.Errorf("failed to aquire resource")
	}

	fmt.Println("success: resource used")
	return nil
}
