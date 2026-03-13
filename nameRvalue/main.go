package main

import "fmt"

func divide(a int, b int) (kushagra int, bharti int) {
	kushagra = a + b
	bharti = a * b
	return
}

func main() {
	fmt.Println("Name return value here we are larning")

	kushagra, bharti := divide(10, 3)
	fmt.Println("kushagra", kushagra, "bharti", bharti)
}
