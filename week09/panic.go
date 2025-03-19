package main

import (
	"fmt"
	"time"
)

func main() {

	m := make(map[string]int)

	go func() {
		m["foo"] = 1
		fmt.Println("First func exiting")
	}()

	go func() {
		m["bar"] = 2
		fmt.Println("Second func exiting")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main exiting")
}
