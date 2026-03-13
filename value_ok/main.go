package main

import "fmt"

func main() {
	fmt.Println("Learning value ok")

	points := map[string]int{
		"a": 10,
		"b": 0, // valid value
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valueC, okC := points["c"]
	fmt.Println("c", valueC, okC)

	if val, ok := points["c"]; ok {
		fmt.Println("c", val)
	} else {
		fmt.Println("c is not present")
	}

	prices := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 30,
	}

	total := 0

	for item, price := range prices {
		fmt.Println(item, "price", price)
		total += price
	}

	fmt.Println("total", total)

	// only keys
	for item := range prices {
		fmt.Println(item)
	}

	// only values
	for _, price := range prices {
		fmt.Println(price)
	}

}
