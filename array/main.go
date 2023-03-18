package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "xyz"

	fmt.Println("Fruit List is : ", fruitList)
	fmt.Println("Fruit List length is : ", len(fruitList))

	var vegList = [3]string{"potato", "peas", "onion"}
	fmt.Println("Vegetable list is : ", vegList)
	fmt.Println("Vegetable list length is : ", len(vegList))
}
