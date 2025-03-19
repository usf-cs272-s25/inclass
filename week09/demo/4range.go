package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("func sleeping")
		time.Sleep(1 * time.Second)
		fmt.Println("func waking up")
		for i := range 4 {
			ch <- i
		}
		close(ch)
	}()

	for a := range ch {
		fmt.Println("answer: ", a)
	}

	fmt.Println("finishing main")
}
