package main

import "fmt"

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	var num int
	fmt.Print("Enter your number: ")
	fmt.Scanln(&num)
	fmt.Println(factorial(num))
}
