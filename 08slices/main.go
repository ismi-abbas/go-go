package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to go slices")
	// didn't declare any value will change array to slice
	var carList = []string{"audi", "ferrari", "myvi"}
	fmt.Printf("Type if fruitlist is %T\n", carList)
	fmt.Println(carList)
	carList = append(carList, "bezza", "axia")
	fmt.Println(carList)
	// [audi ferrari myvi bezza axia]

	carList = append(carList[1:]) // same as slice in javascript that has start and endpoint
	fmt.Println(carList)
	// [ferrari myvi bezza axia]
	carList = append(carList[1:3])
	fmt.Println(carList)
	// [ferrari myvi bezza axia][myvi bezza]

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 923
	highscores[2] = 565
	highscores[3] = 790
	fmt.Println("Value", highscores)
	// Value [234 923 565 790]
	highscores = append(highscores, 666, 867, 321)
	fmt.Println("Value", highscores)
	// Value [234 923 565 790 666 867 321]
	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println("Value", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

}
