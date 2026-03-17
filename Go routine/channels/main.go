package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

// channel is a pipe through which we can send and recive values between go routines
// channels --> pipe -> ek side se data send ek side se recive
// go routine ke bech ka communication ka way hai channel

// send
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number: ", num)
// 		time.Sleep(time.Second * 1)
// 	}

// 	// num := <-numChan
// 	// fmt.Println("Processing number: ", num)
// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult // blocking
// }

// goRoutine synchronizer
// func task(done chan bool) {

// 	defer func() { done <- true }()

// 	fmt.Println("Task is running")

// }

// func emailSender(emailChanel chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChanel {
// 		fmt.Println("Sending email to: ", email)

// 		time.Sleep(time.Second * 1)
// 	}
// }

func main() {
	fmt.Println("Learning about channels")

	// now multiple channel 
	chan1 :=  make (chan int)
	chan2 :=  make (chan string)

	go func(){
		chan1 <-10
	}()

	go func(){
		chan2 <-"pong"
	}()


	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Recived from chan1: ", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Recived from chan2: ", chan2Val)
		}
	}

	// emailChanel := make(chan string, 100) // bufferd channel non-blocking

	// done := make(chan bool)

	// go emailSender(emailChanel, done)

	// for i := 0; i < 10; i++ {
	// 	emailChanel <- fmt.Sprintf("%d@example.com", i)
	// }

	// fmt.Println("done sending")

	// emailChanel <- "1@example.com"
	// emailChanel <- "2@example.com"

	// fmt.Println("receive1: ", <-emailChanel)
	// fmt.Println("receive2: ", <-emailChanel)

	// close(emailChanel)
	// <-done

	// done := make(chan bool) // unbuffered channel

	// go task(done)

	// // wait for the task to complete
	// <-done //block

	// result := make(chan int)

	// go sum(result, 10, 20)

	// res := <-result

	// fmt.Println("Here is the result: ", res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// numChan <- 5

	// time.Sleep(time.Second * 1)

	// messageChan := make(chan string)

	// iska mtlb hm data channel ke andr send kr rhe hai
	// messageChan <- "ping" // blocking

	// iska mtlb hm data channel se recive kr rhe hai
	// message := <-messageChan

	// fmt.Println("Recived message: ", message)

}
