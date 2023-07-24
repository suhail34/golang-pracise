package main

import (
	"fmt"

	patterns "example.com/design-patterns/patterns"
)

func main() {
	done := make(chan struct{})
	for i := 0; i < 30; i++ {
		go func() {
			instance := patterns.GetInstance()
			done <- struct{}{}
			fmt.Printf("Using instance: %p\n", instance)
		}()
		// go patterns.GetInstance(done)
	}

	for i := 0; i < 30; i++ {
		<-done
	}

	car := patterns.GetShop(patterns.CarType)
	fmt.Println(car.GetName())
	mobile := patterns.GetShop(patterns.MobileType)
	fmt.Println(mobile.GetName())
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	// fmt.Scanln()
}
