package main

// go routine basically light weight thread hote hai jn bhi multithreadging
// and concurrnetly chize krni hai

// go keyword :- go keyword se hum kisi bhi function ko go routine mei run krwa sakte hai

import (
	"fmt"
	"sync"
)

// always recive a pointer of wait group
func task(id int, wg *sync.WaitGroup) {
	// defer keyword se hum kisi bhi function ko run krwa sakte hai
	defer wg.Done()

	fmt.Println("Doing Task", id)
}

func main() {
	fmt.Println("Starting learning Go Routine")
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {

		wg.Add(1)
		// non blocking iteration
		go task(i, &wg) // ab ek sth concurnetly run hogi ek sth mei

	}

	// main thread ko rokne ke liye
	wg.Wait()

}
