package main

import (
	"fmt"

	patterns "example.com/design-patterns/patterns"
)

func main() {
	done := make(chan struct{})
	for i := 0; i < 30; i++ {
		go func() {
			//_ = patterns.GetInstanceNew()
      _ = patterns.GetInstance()
			done <- struct{}{}
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
}
