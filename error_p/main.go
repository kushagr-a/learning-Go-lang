package main

import "fmt"

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0,
			fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Here we are learning about error handling in go")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result:", result)
}
