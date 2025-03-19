package main

import (
	"fmt"
	"time"
)

func main() {
	// chan is a "thread-safe queue" used to communicate
	// between goroutines

	// This channel is "unbuffered". Reads of this channel
	// are blocking/synchronous
	ch := make(chan int)

	go func(c chan int) {
		// Write the integer 1 into the channel
		time.Sleep(2 * time.Second)
		c <- 1
	}(ch)

	// Read the channel and put its results into a local var
	// This waits until data is available on the unbuffered channel
	i := <-ch
	fmt.Println("i: ", i)
}
