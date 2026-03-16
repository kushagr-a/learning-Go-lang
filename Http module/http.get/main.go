package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Learning about HTTP GET")

	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
}


//