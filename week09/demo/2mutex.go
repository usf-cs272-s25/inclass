package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	m := map[string]int{}

	var mu sync.Mutex

	go func() {
		mu.Lock()
		defer mu.Unlock()
		m["foo"] = 1
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		m["bar"] = 1
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("finishing main")
}
