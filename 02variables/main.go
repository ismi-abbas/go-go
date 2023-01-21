package main

import "fmt"

// Public varibale will be start as a capital letter
const LoginToken string = "testing778123"

func main() {
	var username string = "Abbas"
	fmt.Println(username)
	fmt.Printf("Variable is type of : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of : %T \n", smallVal)

	var smallFloat float32 = 255.25512321321
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of : %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is type of : %T \n", anotherVariable)

	// implicit type
	var website = "google.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 3000.000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is type of : %T \n", LoginToken)

}
