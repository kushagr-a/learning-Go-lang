package main

import (
	"fmt"
	"go-module/internal/greet"
)

func main() {
	message := greet.Hello("John")
	fmt.Println(message)
}
