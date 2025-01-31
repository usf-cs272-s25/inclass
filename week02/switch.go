package main

import "fmt"

func main() {
	s := "foo"

	switch s {
	case "foo":
		fmt.Println("found a foo")
	case "bar":
		fmt.Println("found a bar")
		fallthrough
	case "baz":
		fmt.Println("found a baz")
	default:
		fmt.Println("not recognized: ", s)
	}
}
