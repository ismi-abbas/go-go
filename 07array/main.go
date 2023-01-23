package main

import "fmt"

func main() {
	fmt.Println("Welcome to go array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is ", fruitList)
	// Fruit list is  [Apple Orange  Banana]
	fmt.Println("Fruit list length is ", len(fruitList))
	// Fruit list is  4

	var vegList [3]string
	vegList[0] = "potato"
	vegList[1] = "cabbage"
	vegList[2] = "salad"
	fmt.Println("Vege list ", vegList)
}
