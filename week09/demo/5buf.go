package main

import (
	"fmt"
	"time"
)

type Person struct {
	First string
	Last  string
}

func main() {

	ch := make(chan Person, 10)
	defer close(ch)

	go func(ch chan Person) {
		fmt.Println("starting long-running goroutine")
		for p := range ch {
			fmt.Println("Person: ", p)
		}
		fmt.Println("channel was closed")
	}(ch)

	ch <- Person{"Sally", "Ride"}
	ch <- Person{"Roger", "Boijoly"}
	ch <- Person{"Christa", "McAuliffe"}

	time.Sleep(100 * time.Millisecond)
}
