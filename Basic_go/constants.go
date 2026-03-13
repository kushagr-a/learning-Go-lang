package main

import "fmt"

func constants() {
	// constant those not change value
	const app = "Go basics"
	fmt.Println(app)

	// typed constant
	const maxUpload int = 20
	fmt.Println(maxUpload)

	// untyped constant
	const pi1 = 3.14159
	fmt.Println(pi1)

	// we can also use multiple constant
	const (
		name = "Kushagra"
		age  = 20
		pi   = 3.14159
	)
	fmt.Println(name, age, pi)
}
