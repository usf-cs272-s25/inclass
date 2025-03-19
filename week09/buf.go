package main

import (
	"fmt"
	"time"
)

func main() {
	// chan is a "thread-safe queue" used to communicate
	// between goroutines

	// This channel is "buffered". Reads of this channel
	// are non-blocking/asynchronous
	ch := make(chan int, 10)

	go func(c chan int) {
		// Write the integer 1 into the channel
		time.Sleep(2 * time.Second)
		c <- 1
		time.Sleep(2 * time.Second)
		c <- 2
		time.Sleep(2 * time.Second)
		c <- 3
		close(c)
		fmt.Println("Anon goroutine exiting")
	}(ch)

	// Range over the buffered channel, removing elements
	// as we go. The range loop stops when the channel gets closed
	// If there's nothing in the channel to read, range will block
	for i := range ch {
		fmt.Println("i: ", i)
	}

	fmt.Println("Main goroutine exiting")
}
