package main

import "fmt"

func main() {
	f := func() {
		i := 0
		fmt.Println("hi there from anon func", i)
		i++
	}

	f()
	f()

	fmt.Println("in main")
}
