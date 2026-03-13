package main

import (
	"fmt"
	"strings"
)

func string_exp() {
	firstname := "kushagra"
	lastname := "bharti"

	fullname := firstname + " " + lastname

	upperFullName := strings.ToUpper(fullname)

	fmt.Println("Fullname :", upperFullName)
}
