package main

import "fmt"

type Loc int

const (
	MyGarage Loc = iota
	ImpoundLot
)

type Car struct {
	Make     string
	Model    string
	Year     int
	Location Loc
}

func Impound(c *Car) {
	c.Location = ImpoundLot
}

func main() {
	var c Car
	c.Make = "Honda"
	c.Model = "Accord"
	c.Year = 2025

	fmt.Println("c: ", c)

	c2 := Car{
		"Chevy", "Camaro", 1965, MyGarage,
	}
	fmt.Println("c2: ", c2)

	cars := []Car{
		{"VW", "Golf", 2024, MyGarage},
		{"Toyota", "GR86", 2023, MyGarage},
		{"Porsche", "959", 1985, MyGarage},
	}

	for idx, val := range cars {
		fmt.Printf("idx: %d, make: %v\n", idx, val)
		Impound(&val)
		fmt.Println("after impound: ", val)
	}

	fmt.Println("Length of cars: ", len(cars))

	for i := 0; i < len(cars); i++ {
		fmt.Println("Car using C-like loop: ", cars[i])
	}
}
