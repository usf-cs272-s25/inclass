package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("foo.txt")
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		return
	}
	defer file.Close()

	b := make([]byte, 1024)
	n, err := file.Read([]byte(b))
	if err != nil {
		fmt.Println("Failed to read file: ", err)
		return
	}
	
	fmt.Println("bytes read: ", n)
	fmt.Println(string(b))
}
