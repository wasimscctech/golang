package main

import "fmt"

/*why struct?
as array and slice holds only single data type what if we have multiple data types values like name, salary, id, and boolean(for yes and no)?
we have struct to hold multiple types of data.
*/

type Employee struct {
	EmpId     int
	EmpName   string
	EmpMobile []string //we have used the camel case for variable to make it global so that these variables can be accessed from the other packages within the same project
}

func main() {
	// emp := Employee{
	// 	226,
	// 	"Wasim",
	// 	[]string{"8668556160", "1234567890"},
	// 	//the values go assigned even if we didn't mention the variables to define the values. it got assigned implicity because the order was same. however if the assignment order is different than the struct variable order then there would be an error like below
	// }
	// fmt.Println(emp)

	// emp2 := Employee{
	// 	226,
	// 	[]string{"8668556160", "1234567890"}, //compiler error
	// 	"Wasim",
	// }

	//we will define the struct variables by using the variable names, it will get assigned correctly doesn't matter what is the order
	emp3 := Employee{
		EmpName:   "Wasim",
		EmpId:     226,
		EmpMobile: []string{"8668556160", "1234567890"},
	}
	fmt.Println(emp3)

	emp4 := emp3 //this will assign the elements of emp3 to emp4
	// emp4 := &emp3 //this will assign the memory location of emp3 to emp4 so if we make any changes in any of those struct the changes will be done for both.
	emp3.EmpName = "Almas"
	fmt.Printf("emp3 is: %v\n", emp3)
	fmt.Printf("emp4 is: %v\n", emp4)

}
