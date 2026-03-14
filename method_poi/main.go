package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("learning method pointer")

	u := User{
		Name: "kushagra",
		Age:  21,
	}

	fmt.Println(u.birthday())
}

func (u *User) birthday() string {
	u.Age++
	u.Name = "kushagra bharti"
	return fmt.Sprintf("my name is %s and i am %d years old", u.Name, u.Age)
}
