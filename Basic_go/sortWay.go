package main

import (
	"fmt"
)

func sortWay() {

	var city string
	city = "LundDon"
	fmt.Println("short way variable", city)

	var name = "sikari"
	fmt.Println("short way variable", name)

	// := declare and assign with type inferance
	city1 := "pglu"
	fmt.Println("short way variable", city1)

	// taking multiple at once
	like, comments := 100, 30
	fmt.Println("Likes and comments", like, comments)
	fmt.Println("Likes", like, "and comments", comments)

}
