package patterns

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleBP struct {
}

var singleBPInstance *singleBP

func GetInstanceNew() *singleBP {
   if singleBPInstance == nil {
     once.Do(
       func() {
         fmt.Println("Creating single Instance Now")
         singleBPInstance = &singleBP{}
       })
   }else {
      fmt.Println("Single Instance Already created")
   }
   return singleBPInstance
}
