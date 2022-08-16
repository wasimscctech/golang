package main

import "fmt"

//call by value
func sum(num int) {
	num = 40
	fmt.Println(num)
}

// call by reference
// func sum(num *int) {
// 	*num = 40
// 	fmt.Println(*num)
// }

func main() {
	num := 4
	sum(num)
	// sum(&num) //call by reference
	fmt.Println(num)
}
