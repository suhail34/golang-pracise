package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Mango"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Kiwi", "Papaya")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 455
	highScore[1] = 415
	highScore[2] = 385
	highScore[3] = 265

	fmt.Println(highScore)
	fmt.Println(len(highScore))

	highScore = append(highScore, 444, 433, 244)

	fmt.Println(highScore)
	fmt.Println(len(highScore))

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	var courses = []string{"angular13", "javascript", "python", "java", "typescript", "c++"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
