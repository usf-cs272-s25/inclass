package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Printf("m: %v\n", m)

	val, ok := m["three"]
	if ok {
		fmt.Println("m[three]: ", val)
	} else {
		fmt.Println("m does not contain the key")
	}

	m["four"] = 4
	if val, ok = m["four"]; ok {
		fmt.Println("val is ", val)
	}
}
