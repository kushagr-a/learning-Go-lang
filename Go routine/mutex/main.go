package main

import (
	"fmt"
	"sync"
	// "time"
)

// Race condition se bachne ke liye mutex use krte hai
// multiple process jb same resource ko modify kr rhe hai
// tb race condition hoti hai

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock() // pura func kbhi lock nhi krna hai jha pe modification ho rhi hai uhi pe krna hai
	p.views++ // iska bottle neck bn rha hai 

}

func main() {
	fmt.Println("Learning about mutex")

	var wg sync.WaitGroup

	myPost := post{views: 0}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println("Final views: ", myPost.views)
}
