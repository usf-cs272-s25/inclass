package main

import (
	"fmt"
	"strings"
)

func main() {
	sl := make([]string, 10)
	sl = append(sl, "foo")
	sl = append(sl, "bar")
	fmt.Println("slice: ", sl)

	data := "this is some text"
	words := strings.Fields(data)
	sl = append(sl, words...)
	fmt.Println("slice: ", sl)
}
