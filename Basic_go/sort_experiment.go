package main

import "fmt"

func sort_experiment() {
	fmt.Println("Sort Experiment")

	items := 2
	pricePerItem := 50

	if total := items * pricePerItem; total > 100 {
		fmt.Println("Total is greater than 100")
	} else {
		fmt.Println("Total is less than 100")
	}

}
