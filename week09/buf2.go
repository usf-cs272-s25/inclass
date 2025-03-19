package main

import (
	"fmt"
	"time"
)

type Person struct {
	First string
	Last  string
}

func foo(c chan Person) {
	// Write the integer 1 into the channel
	time.Sleep(2 * time.Second)
	c <- Person{"Sally", "Ride"}
	time.Sleep(2 * time.Second)
	c <- Person{"Roger", "Boisjoly"}
	time.Sleep(2 * time.Second)
	c <- Person{"Christa", "McAuliffe"}
	close(c)
	fmt.Println("foo goroutine exiting")
}

func main() {
	// chan is a "thread-safe queue" used to communicate
	// between goroutines

	// This channel is "buffered". Reads of this channel
	// are non-blocking/asynchronous
	ch := make(chan Person, 10)

	go foo(ch)

	// Range over the buffered channel, removing elements
	// as we go. The range loop stops when the channel gets closed
	// If there's nothing in the channel to read, range will block
	for p := range ch {
		fmt.Println("p: ", p)
	}

	fmt.Println("Main goroutine exiting")
}
