package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value string
	Error error
}

func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {

	// 4
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Println("image processing: ", url)

	resultChan <- Result{
		Value: url,
		Error: nil,
	}
}

// fan out fan in pattern
// aggregation fan in
// worker pool
// main function running on main thread
func main() {
	fmt.Println("Learning about concurrency patterns")

	// 1
	var wg sync.WaitGroup

	resultChan := make(chan Result, 10)

	startTime := time.Now()

	// 2
	wg.Add(2)
	go worker("image_1.png", &wg, resultChan) // also added here
	go worker("image_2.png", &wg, resultChan)

	// 3
	wg.Wait()

	close(resultChan)

	for result := range resultChan {
		fmt.Println("result: ", result)
	}

	//

	// fmt.Println(result1)
	// fmt.Println(result2)
	// fmt.Println(result3)

	fmt.Println("time taken: ", time.Since(startTime))
}
