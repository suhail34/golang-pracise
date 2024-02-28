package main

import (
	"fmt"

	behavioralpatterns "example.com/design-patterns/behavioral-patterns"
	patterns "example.com/design-patterns/creational-patterns"
)

func main() {
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func() {
			//_ = patterns.GetInstanceNew()
      _ = patterns.GetInstance()
			done <- struct{}{}
		}()
		// go patterns.GetInstance(done)
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	car := patterns.GetShop(patterns.CarType)
	fmt.Println(car.GetName())
	mobile := patterns.GetShop(patterns.MobileType)
	fmt.Println(mobile.GetName())

  ctx := behavioralpatterns.GetContext()
  ctx.GetState()
  ctx.SetState(&behavioralpatterns.Published{})
  ctx.GetState()
}
