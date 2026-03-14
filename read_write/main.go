package main

import (
	"fmt"
	"os"
)

func readFile() {
	read, err := os.ReadFile("read.txt")

	if err != nil {
		fmt.Println("error reading file", err)
		return
	}

	// defer read.close ()

	fmt.Println("file content", string(read))
}

// func writeFile() {
// 	// os.WriteFile only returns an error
// 	err := os.WriteFile("read.txt", []byte("\nHeello from main Go"), 0644)
// 	if err != nil {
// 		fmt.Println("error writing file", err)
// 		return
// 	}
// 	fmt.Println("file written successfully")
// }

func writeFile() {
	file, err := os.OpenFile("read.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("\nHello from Go")
	if err != nil {
		fmt.Println("error writing file", err)
		return
	}

	fmt.Println("data appended successfully")
}

func main() {
	fmt.Println("Learning how to use read and write files")

	// read file
	defer readFile()

	// write file
	writeFile()
}
