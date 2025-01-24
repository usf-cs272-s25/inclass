package main

import "fmt"

type Person struct {
	Name   string
	Height int
}

func PrintPerson(p *Person) {
	fmt.Println("name: ", p.Name)
	fmt.Println("height: ", p.Height)
}

func main() {
	s := []int{}

	s = append(s, 1)
	fmt.Println("s: ", s)

	s = append(s, 2)
	fmt.Println("s: ", s)

	people := []Person{
		{
			"Phil",
			69,
		},
		{
			"Steph Curry",
			77,
		},
	}

	ryan := Person{"Ryan", 72}
	people = append(people, ryan)

	for _, p := range people {
		PrintPerson(&p)
	}
}
