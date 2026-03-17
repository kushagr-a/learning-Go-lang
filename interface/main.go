package main

import "fmt"

// paymenter is an interface.
// Any struct that has the 'pay' method automatically implements this interface.
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter // this field can hold razorpay, stripe, or anything that implements paymenter
}

func (p payment) makePayment(amount float32) {
	// calls the pay method of the specific gateway we stored in the struct
	p.gateway.pay(amount)
}

// razorpay struct
type razorpay struct{}

// razorpay implementation of the interface
func (r razorpay) pay(amount float32) {
	fmt.Println("razorpay payment done", amount)
}

// stripe struct (another payment gateway)
type stripe struct{}

// stripe implementation of the interface
func (s stripe) pay(amount float32) {
	fmt.Println("stripe payment done", amount)
}

func main() {
	// Example 1: Use Razorpay
	razorpayGw := razorpay{}
	newPayment := payment{
		gateway: razorpayGw,
	}
	newPayment.makePayment(100)

	// Example 2: Use Stripe (easy to switch because of the interface)
	stripeGw := stripe{}
	newPayment2 := payment{
		gateway: stripeGw,
	}
	newPayment2.makePayment(200)
}
