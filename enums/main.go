package main

import "fmt"

// enumerated types
// we are implementing enums in go using const
// in go we are making own coustom types like:- type MyType string
// type OrderStatus int
type OrderStatus string

const (
	Received   OrderStatus = "received" // iota basically pre decleard identifier hai for untyped integer
	Confirmed  OrderStatus = "confirmed"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Prepared   OrderStatus = "prepared"
	Delivered  OrderStatus = "delivered"
	Cancelled  OrderStatus = "cancelled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to", status)
}

func main() {
	changeOrderStatus(Cancelled)
}
