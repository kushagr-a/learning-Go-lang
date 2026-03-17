// go routine basically light weight thread hote hai jn bhi multithreadging
// and concurrnetly chize krni hai

// go keyword :- go keyword se hum kisi bhi function ko go routine mei run krwa sakte hai
package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing Task", id)
}

func main() {
	fmt.Println("Starting learning Go Routine")
	for i := 0; i <= 10; i++ {
		// non blocking iteration
		// go task(i) // ab ek sth concurnetly run hogi ek sth mei

		//  here also use the concept of closure
		go func(id int) {
			fmt.Println("Doing Task Inside the func: ", id)
		}(i)
	}

	// main thread ko rokne ke liye
	time.Sleep(time.Second * 2) // aise nhi rokte hai hm 
}
