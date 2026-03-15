package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func main() {

	fmt.Println("Learning about JSON Decoding")

	http.HandleFunc("/test", testHandle)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
