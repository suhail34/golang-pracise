package main

import "fmt"

const Public string = "hello" // Capital letter for declaring a global variable

func main() {
	var username string = "suhail"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.45553423
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type : %T \n", anotherVar)

	fmt.Println(Public)
	fmt.Printf("Variable is of type : %T \n", Public)
}
