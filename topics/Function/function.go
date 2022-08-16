package main

import "fmt"

func show() {
	fmt.Println("this is non parameterized function")
}

// func sum(x, y int) { //we can mention only one data type if all the parameters passed in the function is same
func sum(x int, y int) {
	fmt.Println("this is parameterized func:", x+y)
}

func summation1(vals ...int) {
	sum := 0
	for _, n := range vals {
		sum = sum + n
	}
	fmt.Println("a verdict function without return value", sum)
}

func summation2(vals ...int) int { //we need to specify return type
	sum := 0
	for _, n := range vals {
		sum = sum + n
	}
	return sum
}

func twoReturnValues(vals ...int) (int, int) { //for multiple return values, we need to specify return types in parenthesis
	sum := 0
	prod := 1
	for _, n := range vals {
		sum = sum + n
		prod = prod * n
	}
	return sum, prod
}

func main() {
	show()

	sum(2, 3)

	summation1(2, 3, 4, 5, 6, 7, 8)

	returnedValue := summation2(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("verdict function with return value", returnedValue)

	sum, prod := twoReturnValues(1, 2, 3, 4, 5)
	fmt.Println(sum)
	fmt.Println(prod)
}
