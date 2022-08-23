package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer")
	var ptr *int
	fmt.Println(ptr) //prints <nil>
	fmt.Printf("the type of pointer is: %T", ptr)

	myVar := 7
	var my_ptr = &myVar
	fmt.Println("The value of a pointer", my_ptr)
	fmt.Println("The value of inside the pointer", *my_ptr)

	*my_ptr = *my_ptr + 2
	fmt.Println("the new value of variable", myVar)
}
