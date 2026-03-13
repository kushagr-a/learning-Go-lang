package main

import (
	"fmt"
)

// length = how many element you can have
// capacity = how many element you can store

func main() {
	fmt.Println("Learning len and cap")

	scores := make([]int, 0, 5)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))

	scores = append(scores, 10)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))

	scores = append(scores, 20)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))

	scores = append(scores, 30)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))

	scores = append(scores, 40)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))

	scores = append(scores, 50)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))

	scores = append(scores, 60)
	fmt.Println("scores", scores)
	fmt.Println("length", len(scores))
	fmt.Println("capacity", cap(scores))
}
