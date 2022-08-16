package main

import "fmt"

type Employee struct {
	EmpName string
}

func (emp Employee) test() {
	fmt.Println(emp.EmpName)
}

func main() {
	emp := Employee{"WasimShaikh"}
	emp.test()

}
