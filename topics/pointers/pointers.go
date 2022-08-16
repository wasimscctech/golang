package main

import "fmt"

type Employee struct {
	EmpName string
}

func main() {

	var x int = 90
	var y *int = &x

	fmt.Println(x)
	fmt.Println(y)  //prints the memory location
	fmt.Println(*y) //value

	var emp1 Employee
	var emp2 *Employee
	fmt.Println(emp2) //this will print "<nil>" as the pointer emp2 is not pointing any memory location
	fmt.Println(emp1) //for string: prints {} empty struct. for int{0}

	emp2 = new(Employee)
	fmt.Println(emp2)  //prints for string: &{}, for int:&{0}
	fmt.Println(*emp2) //prints {}

}
