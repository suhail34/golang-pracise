package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")

	total, message := Add(45, 12, 6)
	fmt.Printf("%v %v and the total is : \n", message, total)
}

func Add(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	if total > 100 {
		return total, "Addition is greater than 100"
	}
	return total, "Addition is less than 100"
}
