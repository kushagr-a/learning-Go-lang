package main

import (
	"fmt"
	"package/auth"
)

func main() {
	fmt.Println("learning about packages")
	auth.LoginWithCredentials("kushu", "kush")
}
