package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Learning about file's")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file: ", err)
	// 	return
	// }

	// fileInfo, err := file.Stat()

	// if err != nil {
	// 	fmt.Println("Error getting file info: ", err)
	// 	return
	// }

	// fmt.Println("File name: ", fileInfo.Name())
	// fmt.Println("File size: ", fileInfo.Size())
	// fmt.Println("Is directory: ", fileInfo.IsDir())
	// fmt.Println("File mode: ", fileInfo.Mode())
	// fmt.Println("File mod time: ", fileInfo.ModTime())

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file: ", err)
	// 	return
	// }
	// defer file.Close()

	// buff := make([]byte, 1024) // array of bytes
	// d, err := file.Read(buff)

	// if err != nil {
	// 	fmt.Println("Error reading file: ", err)
	// 	return
	// }

	// println("data", d, string(buff))

	// defer file.Close()

	// file, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	fmt.Println("Error reading file: ", err)
	// 	return
	// }
	// println("==>", string(file))

	// read folder
	// dir, err := os.Open("../")

	// if err != nil {
	// 	fmt.Println("Error reading folder: ", err)
	// 	return
	// }

	// defer dir.Close()

	// fileInfos, err := dir.Readdir(-1)

	// if err != nil {
	// 	fmt.Println("Error reading folder: ", err)
	// 	return
	// }

	// for _, fi := range fileInfos {
	// 	fmt.Println("Folder name: ", fi.Name())
	// }

	// create a file
	// file, err := os.Create("example2.txt")

	// if err != nil {
	// 	fmt.Println("Error creating file: ", err)
	// 	return
	// }

	// defer file.Close()

	// file.WriteString("hi go land") // append krta hai replace nhi

	// getting from one file and transforming to another file (streming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file: ", err)
	// 	return
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file: ", err)
	// 	return
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		break
	// 	}
	// 	writer.WriteString(line)
	// }

	// writer.Flush()

	// delete a file

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file: ", err)
	// 	return
	// }
	// defer sourceFile.Close()

	err := os.Remove("example2.txt")
	if err != nil {
		fmt.Println("Error removing file: ", err)
		return
	}

}
