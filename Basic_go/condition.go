package main

import "fmt"

func condition() {
	fmt.Println("Print  From Condition Block")

	score := 90
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

}

