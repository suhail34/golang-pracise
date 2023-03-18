package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}() // fork point

	<-done // join point
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff from child")
}
