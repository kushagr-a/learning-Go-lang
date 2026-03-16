// json.Unmarshal in Go converts JSON-encoded data ([]byte) into Go data structures (structs, maps, or slices) using the encoding/json package

package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
}

func main() {
    jsonData := []byte(`{"name": "Alice", "age": 30}`)
    var p Person

    err := json.Unmarshal(jsonData, &p)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(p.Name, p.Age) // Output: Alice 30
}
