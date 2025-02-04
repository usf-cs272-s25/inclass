package main

import "fmt"

type MotorVehicle interface {
	WheelCount() int
}

type Car struct {
}

func (c Car) WheelCount() int {
	return 4
}

type Motorcycle struct {
}

func (mc Motorcycle) WheelCount() int {
	return 2
}

func PrintWheelCount(m MotorVehicle) {
	fmt.Println("Wheel count: ", m.WheelCount())
}

func main() {
	var c Car
	// fmt.Println("Car wheel count: ", c.WheelCount())
	PrintWheelCount(c)

	var mc Motorcycle
	PrintWheelCount(mc)
}
