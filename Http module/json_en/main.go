package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "success",
		"dateTime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)

}

func main() {
	fmt.Println("Learning about JSON Encoding")

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
