package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Error handling in go")

	// go dont use exceptions for normal failures
	// function -> return error as normal return value

	// value, error := something()

	err := run()
	if err != nil {
		fmt.Println("error:", err)
	}

}

func run() error {
	input := "8"
	level, err := parseLevel(input)
	if err != nil {
		return err
	}
	fmt.Println("level:", level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value, error)
	//  nil error -> success
	// non nil -> failure

	// pattern

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("level must be  a  number")
	}
 
	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must be between 1 and 5")
	}
	return n, nil
}
