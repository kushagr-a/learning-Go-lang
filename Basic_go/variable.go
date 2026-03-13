package main

import "fmt"

func variable() {
	// var name type
	var channelName string
	channelName = "Kushagra"

	var year int = 2026

	var rating float64 = 4.9
	// when you need more accurate then use float64
	// when you need less accurate then use float32

	fmt.Println("Channel name", channelName)
	fmt.Println("year name", year)
	fmt.Println("rating name", rating)

}
