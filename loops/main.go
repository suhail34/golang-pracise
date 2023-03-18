package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println()

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println()

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto skna
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

skna:
	fmt.Println("Jumping at suhail khan")
}
