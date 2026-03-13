package main

import "fmt"

func main() {
	fmt.Println("Enter a number between 1 and 7:")
	var num int    // assign a value to num
	fmt.Scan(&num) // read a value from the user

	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input")
	}
}
