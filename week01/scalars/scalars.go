package main

import "fmt"

func doSomething(i int) (bool, error) {
	if i == 1 {
		fmt.Println("i is 1!")
	} else {
		fmt.Println("i is not 1!")
	}
	return true, nil
}

func main() {
	var i int
	fmt.Println("i: ", i)

	b, err := doSomething(1)
	if err == nil {
		//use b
	} else {
		// b is invalid
	}
	fmt.Println("b: ", b)

}
