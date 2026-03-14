package main

import "fmt"

// handalling for data or model particular represent the model or need struct
// struct groups related feilds into one type

// we are using struct for creating user defined data type
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("learning struct")

	user1 := User{
		ID:    1,
		Name:  "John",
		Email: "[EMAIL_ADDRESS]",
		Age:   25,
	}

	// struct is mutable by default
	user1.Name = "kushu"

	// stuct creat with partial feilds
	user2 := User{
		Name:  "John",
		Email: "[EMAIL_ADDRESS]",
	}

	fmt.Println(user1)
	fmt.Println("partial struct", user2)
}
