package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointer")
	// init integer pointer
	var ptr *int
	fmt.Println("Value if pointer is ", ptr)
	// Value if pointer is  <nil>

	myNumber := 25
	// create pointer point to myNumber
	var ptr_2 = &myNumber
	fmt.Println("Value if pointer is ", ptr_2)
	// Value if pointer is  0xc00000e0b8
	fmt.Println("Value if pointer is ", *ptr_2)
	// Value if pointer is  25

	*ptr_2 = *ptr_2 * 2
	fmt.Println("New value is :", myNumber)
}
