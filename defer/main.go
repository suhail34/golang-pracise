package main

import "fmt"

func main() {
	fmt.Println("Defer in GoLang")
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
	Defer()
}

func Defer() {
	defer fmt.Println("Hello World")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v\n", i)
	}
}
