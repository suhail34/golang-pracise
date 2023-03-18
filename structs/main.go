package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	suhail := User{"Suhail", "skna@gmail.com", true, 21}
	fmt.Println(suhail)
	fmt.Printf("Suhail details are: %+v\n", suhail)
	fmt.Printf("Name is %v: and email is: %v\n", suhail.Name, suhail.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
