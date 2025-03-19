package main

import (
	"fmt"
	"time"
)

func main() {

	m := map[string]int{}

	go func() {
		m["foo"] = 1
	}()

	go func() {
		m["bar"] = 1
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("finishing main")
}
