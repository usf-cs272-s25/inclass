package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		"Phil",
		57,
	}

	var p2 Person
	p2.Name = "Manasa"
	p2.Age = 20

	fmt.Println("p: ", p)

	fmt.Println("p2: ", p2)
}
