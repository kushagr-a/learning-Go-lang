package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("learning method and value")

	u := User{
		Name: "John",
		Age:  25,
	}

	fmt.Println(u.Intro())
}

// value reciver means this method receives a copy of the user
func (u User) Intro() string {
	return fmt.Sprintf("my name is %s and i am %d years old", u.Name, u.Age)
}
