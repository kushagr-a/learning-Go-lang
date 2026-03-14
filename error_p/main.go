package main

import "fmt"

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0,
			fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func add(a int, b int) (int, error) {
	if a < 0 && b > 5 {
		return 0,
			fmt.Errorf("Number must be greater than 0 and less than 5")
	}
	return a + b, nil
}

func main() {
	fmt.Println("Here we are learning about error handling in go")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result:", result)

	resultSum, err := add(2, 3)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("resultSum:", resultSum)
}
