package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Methods in Golang")
	emp := Employee{1, "suhail", 21, time.Now()}
	emp.GetEmpDetails()
	fmt.Println(emp.SetEmpDetails(emp.Joined).Format("01-02-2006 15:04:05 Monday"))
	emp.GetEmpDetails()
}

type Employee struct {
	id     int
	Name   string
	Age    int
	Joined time.Time
}

func (emp Employee) GetEmpDetails() {
	fmt.Printf("Employee Name: %v Employee Age: %v Employee Time: %v\n", emp.Name, emp.Age, emp.Joined.Format("01-02-2006 15:04:05 Monday"))
}

func (emp *Employee) SetEmpDetails(joined time.Time) time.Time {
	emp.Joined = time.Time.AddDate(joined, 1, 2, 3)
	return emp.Joined
}
