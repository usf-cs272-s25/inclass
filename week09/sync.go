package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Map is shared between two anon goroutines
	m := make(map[string]int)

	// "Mutual Exclusion" object is used to synchronize access
	// to shared memory
	var mu sync.Mutex

	go func() {
		mu.Lock()
		// defer is idiomatic Go because it guards against
		// early function returns better than calling Unlock at the end
		defer mu.Unlock()
		m["foo"] = 1
		fmt.Println("First func exiting")
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		m["bar"] = 2
		fmt.Println("Second func exiting")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main exiting")
}
