package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 1
var mutex sync.Mutex

func increment() {
  mutex.Lock()
  defer mutex.Unlock()
  current := counter
  current = 2*current
  time.Sleep(time.Millisecond)
  counter = current
  fmt.Println("Counter : ", counter)
}

func main() {
  for i:=0; i<5; i++ {
    go increment()
  }
  time.Sleep(time.Second)
  fmt.Println(counter)
}
