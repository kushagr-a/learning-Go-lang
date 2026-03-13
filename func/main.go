package main

import "fmt"

// add function
func add(a int, b int) int {
	return a + b
}

// multiple return values
func swap(a int, b int) (int, int) {
	return b, a
}

// return type with multiple returns
func SumandProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	fmt.Println("Learning func")

	result := add(10, 20)
	fmt.Println("result", result)

	swap1, swap2 := swap(10, 20)
	fmt.Println("swap", swap1, swap2)

	sum, product := SumandProduct(10, 20)
	fmt.Println("sum", sum, "product", product)

	// some times not want to return something
	onlySum, _ := SumandProduct(10, 20)
	fmt.Println("onlySum", onlySum)

	// named return values
	
}
