package main

import "fmt"

func main() {
	fmt.Println("Hello GO")
	for i := 0; i < 2; i++ {
		fmt.Println("Hello $i")
	}
}

// GO variables
// Some place can be put semicolon, some are not

// Types
// case sensitive
// public function will be exported in capital letter
// Define the types in advance
// Everything in GO is a type
// Types - String, Bool, Interger(uint8, uint64), Floating, Complex, Array, Slices, Maps, Structs, Pointers, Functions, Channels, ....
