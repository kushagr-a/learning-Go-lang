package main

import "fmt"

func main() {
	fmt.Println("Learning map")

	// map is a collection of key-value pairs
	// map[keyType]valueType

	age := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	fmt.Println(age)

	// access the value
	fmt.Println(age["Alice"])

	// add new key-value pair
	age["David"] = 40
	fmt.Println(age)

	// update the value
	age["Alice"] = 26
	fmt.Println(age)

	// delete the key-value pair
	delete(age, "Bob")
	fmt.Println(age)

	// make map([keyType]valueType)
	var scores map[string]int
	fmt.Println(scores, scores["Alice"])

	
}
