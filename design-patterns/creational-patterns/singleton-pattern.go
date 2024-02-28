package creationalpatterns

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Single Instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created.")
		}
	} else {
		fmt.Println("Single Instance already created.")
	}
	// done <- struct{}{}
	return singleInstance
}
