package main

import (
	"fmt"
)

func int_exp() {
	views := 1000
	views2 := 5000
	totalViews := views + views2

	fmt.Println("Total Views :", totalViews)

	likes :=10
	likes++
	likes ++

	avgViews := totalViews / likes
	fmt.Println("Average Views per Like :",totalViews, likes ,avgViews)

}
	 